package venus

import (
	"context"
	"time"

	"github.com/no-mole/venus/proto/pbnamespace"

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
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx) //must admin
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	if !isAdmin {
		return &pbuser.UserInfo{}, errors.ErrorGrpcPermissionDenied
	}
	if info.Role != pbuser.UserRole_UserRoleAdministrator.String() && info.Role != pbuser.UserRole_UserRoleMember.String() {
		info.Role = pbuser.UserRole_UserRoleMember.String()
	}
	if info.Password == "" {
		info.Password = structs.DefaultPassword
		info.ChangePasswordStatus = pbuser.ChangePasswordStatus_ChangePasswordStatusNo
	}
	err = validate.Validate.Struct(info)
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	return s.server.UserRegister(ctx, info)
}

func (s *Server) UserUnregister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	if !isAdmin {
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
		return &pbuser.LoginResponse{}, errors.ErrorGrpcUserNotExistOrPasswordNotMatch
	}
	if info.Status == pbuser.UserStatus_UserStatusDisable {
		return &pbuser.LoginResponse{}, errors.ErrorGrpcPermissionDenied
	}
	return s.genUserLoginResponse(ctx, info)
}

func (s *Server) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !isAdmin {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.UserChangeStatus(ctx, req)
}

func (s *Server) UserList(ctx context.Context, _ *emptypb.Empty) (*pbuser.UserListResponse, error) {
	resp := &pbuser.UserListResponse{}
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	if !isAdmin {
		return resp, errors.ErrorPermissionDenied
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

func (s *Server) UserNamespaceList(ctx context.Context, req *pbuser.UserNamespaceListRequest) (*pbnamespace.NamespaceUserListResponse, error) {
	resp := &pbnamespace.NamespaceUserListResponse{}
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx) //must admin
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	if isAdmin {
		allNamespaces, err := s.NamespacesList(ctx, nil)
		if err != nil {
			return resp, err
		}
		resp.Items = make([]*pbnamespace.NamespaceUserInfo, 0, len(allNamespaces.Items))
		for _, item := range allNamespaces.Items {
			resp.Items = append(resp.Items, &pbnamespace.NamespaceUserInfo{
				NamespaceUid:   item.NamespaceUid,
				NamespaceAlias: item.NamespaceAlias,
				Role:           string(auth.PermissionWriteRead),
			})
		}
	} else {
		err = s.state.NestedBucketScan(ctx, [][]byte{
			[]byte(structs.UserNamespacesBucketName),
			[]byte(req.Uid),
		}, func(k, v []byte) error {
			item := &pbnamespace.NamespaceUserInfo{}
			err := codec.Decode(v, item)
			if err != nil {
				return err
			}
			resp.Items = append(resp.Items, item)
			return nil
		})
	}
	return resp, errors.ToGrpcError(err)
}

func (s *Server) UserDetails(ctx context.Context, _ *emptypb.Empty) (*pbuser.LoginResponse, error) {
	claims, login := auth.FromContextClaims(ctx)
	if !login {
		return &pbuser.LoginResponse{}, errors.ErrorGrpcNotLogin
	}
	info, err := s.UserLoad(ctx, claims.UniqueID)
	if err != nil {
		return &pbuser.LoginResponse{}, errors.ToGrpcError(err)
	}
	return s.genUserLoginResponse(ctx, info)
}

func (s *Server) genUserLoginResponse(ctx context.Context, info *pbuser.UserInfo) (*pbuser.LoginResponse, error) {
	var roles map[string]auth.Permission

	tokenType := auth.TokenTypeUser
	if info.Role == pbuser.UserRole_UserRoleAdministrator.String() {
		tokenType = auth.TokenTypeAdministrator
	}
	//先生成token以获取UserNamespaceList
	token := auth.NewJwtTokenWithClaim(time.Now().Add(s.config.TokenTimeout), info.Uid, info.Name, tokenType, roles)
	ctx = auth.WithContext(ctx, token)

	resp, err := s.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{Uid: info.Uid})
	if err != nil {
		return &pbuser.LoginResponse{}, err
	}

	//admin 不需要把namespace 信息写入token中
	if !(info.Role == pbuser.UserRole_UserRoleAdministrator.String()) {
		roles = make(map[string]auth.Permission, len(resp.Items))
		for _, item := range resp.Items {
			roles[item.NamespaceUid] = auth.Permission(item.Role)
		}
	}

	token = auth.NewJwtTokenWithClaim(time.Now().Add(s.config.TokenTimeout), info.Uid, info.Name, tokenType, roles)
	tokenString, err := s.authenticator.Sign(ctx, token)
	if err != nil {
		return &pbuser.LoginResponse{}, errors.ToGrpcError(err)
	}
	return &pbuser.LoginResponse{
		Uid:                  info.Uid,
		Name:                 info.Name,
		Role:                 info.Role,
		AccessToken:          tokenString,
		TokenType:            "Bearer",
		NamespaceItems:       resp.Items,
		ChangePasswordStatus: info.ChangePasswordStatus,
	}, nil
}

func (s *Server) UserChangePassword(ctx context.Context, req *pbuser.ChangePasswordRequest) (*pbuser.UserInfo, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	return s.server.UserChangePassword(ctx, req)
}

func (s *Server) UserResetPassword(ctx context.Context, req *pbuser.ResetPasswordRequest) (*pbuser.UserInfo, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx) //must admin
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	if !isAdmin {
		return &pbuser.UserInfo{}, errors.ErrorGrpcPermissionDenied
	}
	err = validate.Validate.Struct(req)
	if err != nil {
		return &pbuser.UserInfo{}, errors.ToGrpcError(err)
	}
	return s.server.UserResetPassword(ctx, req)
}

func (s *Server) UserSync(info *pbuser.UserInfo) (*pbuser.LoginResponse, error) {
	if info.Uid == "" {
		return nil, errors.ErrorGrpcNotLogin
	}
	userInfo, err := s.UserLoad(s.ctx, info.Uid)
	if err != nil {
		if err != errors.ErrorUserNotExist {
			return nil, errors.ToGrpcError(err)
		}
		userInfo, err = s.UserRegister(s.ctx, &pbuser.UserInfo{
			Uid:     info.Uid,
			Name:    info.Name,
			Updater: "venus",
		})
		if err != nil {
			return nil, err
		}
	}
	return s.genUserLoginResponse(s.ctx, userInfo)
}
