package local

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/venus/auth"

	"github.com/no-mole/venus/agent/venus/secret"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) UserRegister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &pbuser.UserInfo{}, errors.ErrorGrpcNotLogin
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	info.Status = pbuser.UserStatus_UserStatusEnable
	info.Password = secret.Confusion(info.Uid, info.Password)
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return info, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return info, f.Error()
	}
	info.Password = ""
	return info, nil
}

func (l *Local) UserUnregister(_ context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	data, err := codec.Encode(structs.UserUnregisterRequestType, info)
	if err != nil {
		return info, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return info, f.Error()
	}
	return info, nil
}

func (l *Local) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	info, err := l.UserLoad(ctx, req.Uid)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	info.Status = req.GetStatus()
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, err
}

func (l *Local) UserLoad(ctx context.Context, uid string) (*pbuser.UserInfo, error) {
	info := &pbuser.UserInfo{}
	data, err := l.fsm.State().Get(ctx, []byte(structs.UsersBucketName), []byte(uid))
	if err != nil {
		return info, err
	}
	err = codec.Decode(data, info)
	if err != nil {
		return info, err
	}
	if info.Uid == "" {
		return info, errors.ErrorUserNotExist
	}
	return info, nil
}

func (l *Local) UserChangePassword(ctx context.Context, req *pbuser.ChangePasswordRequest) (*pbuser.UserInfo, error) {
	info, err := l.UserLoad(ctx, req.Uid)
	if err != nil {
		return &pbuser.UserInfo{}, err
	}
	oldPassword := secret.Confusion(info.Uid, req.OldPassword)
	if oldPassword != info.Password {
		return &pbuser.UserInfo{}, errors.ErrorGrpcUserNotExistOrPasswordNotMatch
	}
	info.Updater = req.Uid
	info.UpdateTime = time.Now().Format(timeFormat)
	info.Password = secret.Confusion(info.Uid, req.NewPassword)
	info.ChangePasswordStatus = pbuser.ChangePasswordStatus_ChangePasswordStatusYes
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return &pbuser.UserInfo{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &pbuser.UserInfo{}, f.Error()
	}
	info.Password = ""
	return info, nil
}

func (l *Local) UserResetPassword(ctx context.Context, req *pbuser.ResetPasswordRequest) (*pbuser.UserInfo, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &pbuser.UserInfo{}, errors.ErrorGrpcNotLogin
	}
	if claims.TokenType != auth.TokenTypeAdministrator {
		return &pbuser.UserInfo{}, errors.ErrorGrpcPermissionDenied
	}
	info, err := l.UserLoad(ctx, req.Uid)
	if err != nil {
		return &pbuser.UserInfo{}, err
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	info.Password = secret.Confusion(info.Uid, structs.DefaultPassword)
	info.ChangePasswordStatus = pbuser.ChangePasswordStatus_ChangePasswordStatusNo
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return &pbuser.UserInfo{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &pbuser.UserInfo{}, f.Error()
	}
	info.Password = ""
	return info, nil
}
