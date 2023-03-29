package server

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/utils"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/proto/pbclient"
	"google.golang.org/grpc/peer"
	"time"
)

const timeFormat = time.RFC3339

func GetClientInfo(ctx context.Context) (*pbclient.ClientInfo, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &pbclient.ClientInfo{}, errors.ErrorNotLogin
	}
	ip := "unknown"
	p, ok := peer.FromContext(ctx)
	if ok {
		ip = p.Addr.String()
	}
	clientHostname, clientIp := utils.FromContext(ctx)
	if clientIp == "" {
		clientIp = "unknown"
	}
	ip += "/" + clientIp
	clientInfo := &pbclient.ClientInfo{
		RegisterTime:      time.Now().Format(timeFormat),
		RegisterAccessKey: fmt.Sprintf("%s(%s)", claims.Name, claims.UniqueID),
		RegisterIp:        ip,
		RegisterHost:      clientHostname,
	}
	return clientInfo, nil
}
