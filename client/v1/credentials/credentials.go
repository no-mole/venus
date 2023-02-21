package credentials

import (
	"context"
	"fmt"
	grpccredentials "google.golang.org/grpc/credentials"
	"sync"
	"time"
)

// Bundle defines gRPC credential interface.
type Bundle interface {
	grpccredentials.Bundle
	UpdateAuthToken(tokenType string, token string, expiredIn time.Duration)
	ShouldUpdateToken() bool
}

// NewBundle constructs a new gRPC credential bundle.
func NewBundle() Bundle {
	return &bundle{
		rc: newPerRPCCredential(),
	}
}

// bundle implements "grpccredentials.Bundle" interface.
type bundle struct {
	rc *perRPCCredential
}

func (b *bundle) ShouldUpdateToken() bool {
	return b.rc.ShouldUpdateToken()
}

func (b *bundle) TransportCredentials() grpccredentials.TransportCredentials {
	//todo
	return nil
}

func (b *bundle) PerRPCCredentials() grpccredentials.PerRPCCredentials {
	return b.rc
}

func (b *bundle) NewWithMode(mode string) (grpccredentials.Bundle, error) {
	// no-op
	return nil, nil
}

func (b *bundle) UpdateAuthToken(tokenType, token string, expiredIn time.Duration) {
	if b.rc == nil {
		return
	}
	b.rc.UpdateAuthToken(tokenType, token, expiredIn)
}

// perRPCCredential implements "grpccredentials.PerRPCCredentials" interface.
type perRPCCredential struct {
	tokenType   string
	authToken   string
	expiredIn   time.Duration
	updateTime  time.Time
	authTokenMu sync.RWMutex
}

func newPerRPCCredential() *perRPCCredential {
	return &perRPCCredential{}
}

func (rc *perRPCCredential) RequireTransportSecurity() bool { return false }

func (rc *perRPCCredential) GetRequestMetadata(ctx context.Context, s ...string) (map[string]string, error) {
	rc.authTokenMu.RLock()
	defer rc.authTokenMu.RUnlock()
	if rc.authToken == "" {
		return nil, nil
	}
	return map[string]string{"authorization": fmt.Sprintf("%s %s", rc.tokenType, rc.authToken)}, nil
}

func (rc *perRPCCredential) UpdateAuthToken(tokenType, token string, expiredIn time.Duration) {
	rc.authTokenMu.Lock()
	defer rc.authTokenMu.Unlock()
	rc.tokenType = tokenType
	rc.authToken = token
	rc.expiredIn = expiredIn
	rc.updateTime = time.Now()
}

func (rc *perRPCCredential) ShouldUpdateToken() bool {
	rc.authTokenMu.RLock()
	defer rc.authTokenMu.RUnlock()
	if rc.expiredIn == 0 {
		return false
	}
	if time.Since(rc.updateTime) < (rc.expiredIn - 5*time.Second) {
		return false
	}
	return true
}
