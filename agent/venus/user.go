package venus

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/secret"

	"github.com/no-mole/venus/agent/venus/validate"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UserRegister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	writable, err := s.authenticator.WritableContext(ctx, "") //must admin
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &pbuser.UserInfo{}, errors.ErrorGrpcPermissionDenied
	}
	if info.Role != pbuser.UserRole_UserRoleAdministrator.String() || info.Role != pbuser.UserRole_UserRoleMember.String() {
		info.Role = pbuser.UserRole_UserRoleMember.String()
	}
	err = validate.Validate.Struct(info)
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	return s.server.UserRegister(ctx, info)
}

func (s *Server) UserUnregister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	writable, err := s.authenticator.WritableContext(ctx, "") //must admin
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &pbuser.UserInfo{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.UserUnregister(ctx, info)
}

func (s *Server) UserLogin(ctx context.Context, req *pbuser.LoginRequest) (*pbuser.LoginResponse, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbuser.LoginResponse{}, errors.ToGrpcError(err)
	}
	info, err := s.UserLoad(ctx, req.Uid)
	if err != nil {
		return &pbuser.LoginResponse{}, errors.ToGrpcError(err)
	}
	if secret.Confusion(req.Uid, req.Password) != info.Password {
		return &pbuser.LoginResponse{}, errors.ErrorUserNotExistOrPasswordNotMatch
	}
	resp, err := s.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{Uid: info.Uid})
	if err != nil {
		return &pbuser.LoginResponse{}, err
	}
	roles := make(map[string]auth.Permission, len(resp.Items))
	for _, item := range resp.Items {
		roles[item.Namespace] = auth.Permission(item.Role)
	}
	token := auth.NewJwtTokenWithClaim(time.Now().Add(s.config.TokenTimeout), info.Uid, info.Name, auth.TokenTypeUser, roles)
	tokenString, err := s.authenticator.Sign(ctx, token)
	if err != nil {
		return &pbuser.LoginResponse{}, errors.ToGrpcError(err)
	}
	return &pbuser.LoginResponse{
		Uid:            info.Uid,
		Name:           info.Name,
		Role:           info.Role,
		AccessToken:    tokenString,
		TokenType:      "Bearer",
		NamespaceItems: resp.Items,
	}, errors.ToGrpcError(err)
}

func (s *Server) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, "") //must admin
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.UserChangeStatus(ctx, req)
}

func (s *Server) UserList(ctx context.Context, _ *emptypb.Empty) (*pbuser.UserListResponse, error) {
	_, err := s.authenticator.WritableContext(ctx, "") //must admin
	resp := &pbuser.UserListResponse{}
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.Scan(ctx, []byte(structs.UsersBucketName), func(k, v []byte) error {
		item := &pbuser.UserInfo{}
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

func (s *Server) UserLoad(ctx context.Context, uid string) (*pbuser.UserInfo, error) {
	info := &pbuser.UserInfo{}
	data, err := s.state.Get(ctx, []byte(structs.UsersBucketName), []byte(uid))
	if err != nil {
		return info, errors.ToGrpcError(err)
	}
	err = codec.Decode(data, info)
	if err != nil {
		return info, errors.ToGrpcError(err)
	}
	if info.Uid == "" {
		return info, errors.ErrorUserNotExist
	}
	return info, nil
}

func (s *Server) UserAddNamespace(ctx context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, info.Namespace)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.UserAddNamespace(ctx, info)
}

func (s *Server) UserDelNamespace(ctx context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, info.Namespace) //must admin
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.UserDelNamespace(ctx, info)
}

func (s *Server) UserNamespaceList(ctx context.Context, req *pbuser.UserNamespaceListRequest) (*pbuser.UserNamespaceListResponse, error) {
	resp := &pbuser.UserNamespaceListResponse{}
	err := s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.UserNamespacesBucketName),
		[]byte(req.Uid),
	}, func(k, v []byte) error {
		item := &pbuser.UserNamespaceInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
