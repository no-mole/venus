package venus

import (
	"context"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/secret"
	"time"

	"github.com/no-mole/venus/proto/pbaccesskey"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/validate"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AccessKeyGen(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*pbaccesskey.AccessKeyInfo, error) {
	return s.remote.AccessKeyGen(ctx, info)
}

func (s *Server) AccessKeyDel(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.AccessKeyDel(ctx, info)
}

func (s *Server) AccessKeyLogin(ctx context.Context, req *pbaccesskey.AccessKeyLoginRequest) (*pbaccesskey.AccessKeyLoginResponse, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbaccesskey.AccessKeyLoginResponse{}, errors.ToGrpcError(err)
	}
	info, err := s.AccessKeyLoad(ctx, req.Ak)
	if err != nil {
		return &pbaccesskey.AccessKeyLoginResponse{}, errors.ToGrpcError(err)
	}
	if secret.Confusion(req.Ak, req.Password) != info.Password {
		return &pbaccesskey.AccessKeyLoginResponse{}, errors.ErrorAccessKeyNotExistOrPasswordNotMatch
	}
	resp, err := s.AccessKeyNamespaceList(ctx, &pbaccesskey.AccessKeyNamespaceListRequest{Ak: info.Ak})
	if err != nil {
		return &pbaccesskey.AccessKeyLoginResponse{}, err
	}
	roles := make(map[string]auth.Permission, len(resp.Items))
	for _, item := range resp.Items {
		roles[item.Namespace] = auth.PermissionReadOnly
	}
	token := auth.NewJwtTokenWithClaim(time.Now().Add(s.config.TokenTimeout), auth.TokenTypeAccessKey, roles)
	tokenString, err := s.authenticator.Sign(ctx, token)
	if err != nil {
		return &pbaccesskey.AccessKeyLoginResponse{}, errors.ToGrpcError(err)
	}

	return &pbaccesskey.AccessKeyLoginResponse{
		Ak:          info.Ak,
		Alias:       info.Alias,
		AccessToken: tokenString,
		TokenType:   "Bearer",
	}, errors.ToGrpcError(err)
}

func (s *Server) AccessKeyChangeStatus(ctx context.Context, req *pbaccesskey.AccessKeyStatusChangeRequest) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.AccessKeyChangeStatus(ctx, req)
}

func (s *Server) AccessKeyList(ctx context.Context, _ *emptypb.Empty) (*pbaccesskey.AccessKeyListResponse, error) {
	resp := &pbaccesskey.AccessKeyListResponse{}
	err := s.state.Scan(ctx, []byte(structs.AccessKeysBucketName), func(k, v []byte) error {
		item := &pbaccesskey.AccessKeyInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		item.Password = ""
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) AccessKeyLoad(ctx context.Context, ak string) (*pbaccesskey.AccessKeyInfo, error) {
	info := &pbaccesskey.AccessKeyInfo{}
	data, err := s.state.Get(ctx, []byte(structs.AccessKeysBucketName), []byte(ak))
	if err != nil {
		return info, errors.ToGrpcError(err)
	}
	err = codec.Decode(data, info)
	if err != nil {
		return info, errors.ToGrpcError(err)
	}
	if info.Ak == "" {
		return info, errors.ErrorAccessKeyNotExist
	}
	return info, nil
}

func (s *Server) AccessKeyAddNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	return s.remote.AccessKeyAddNamespace(ctx, info)
}

func (s *Server) AccessKeyDelNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	return s.remote.AccessKeyDelNamespace(ctx, info)
}

func (s *Server) AccessKeyNamespaceList(ctx context.Context, req *pbaccesskey.AccessKeyNamespaceListRequest) (*pbaccesskey.AccessKeyNamespaceListResponse, error) {
	resp := &pbaccesskey.AccessKeyNamespaceListResponse{}
	err := s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.AccessKeyNamespacesBucketName),
		[]byte(req.Ak),
	}, func(k, v []byte) error {
		item := &pbaccesskey.AccessKeyNamespaceInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
