package local

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"time"

	"github.com/no-mole/venus/agent/venus/auth"

	"github.com/no-mole/venus/agent/venus/secret"

	"github.com/no-mole/venus/proto/pbaccesskey"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) AccessKeyGen(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*pbaccesskey.AccessKeyInfo, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &pbaccesskey.AccessKeyInfo{}, errors.ErrorGrpcNotLogin
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	info.Status = pbaccesskey.AccessKeyStatus_AccessKeyStatusEnable
	ak := md5.Sum([]byte(time.Now().String()))
	info.Ak = base64.RawURLEncoding.EncodeToString(ak[:])
	pwd := md5.Sum(ak[:])
	pwdOrigin := base64.RawURLEncoding.EncodeToString(pwd[:])

	info.Password = secret.Confusion(info.Ak, pwdOrigin)

	data, err := codec.Encode(structs.AccessKeyGenRequestType, info)
	if err != nil {
		return info, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &pbaccesskey.AccessKeyInfo{}, f.Error()
	}
	info.Password = pwdOrigin
	return info, nil
}

func (l *Local) AccessKeyDel(_ context.Context, info *pbaccesskey.AccessKeyDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.AccessKeyDelRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) AccessKeyChangeStatus(ctx context.Context, req *pbaccesskey.AccessKeyStatusChangeRequest) (*emptypb.Empty, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	info, err := l.AccessKeyLoad(ctx, req.Ak)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	info.Status = req.GetStatus()
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	data, err := codec.Encode(structs.AccessKeyGenRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, err
}

func (l *Local) AccessKeyLoad(ctx context.Context, ak string) (*pbaccesskey.AccessKeyInfo, error) {
	info := &pbaccesskey.AccessKeyInfo{}
	data, err := l.fsm.State().Get(ctx, []byte(structs.AccessKeysBucketName), []byte(ak))
	if err != nil {
		return info, err
	}
	err = codec.Decode(data, info)
	if err != nil {
		return info, err
	}
	if info.Ak == "" {
		return info, errors.ErrorAccessKeyNotExist
	}
	return info, nil
}
