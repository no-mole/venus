package venus

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UserRegister(_ context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	info.Status = pbuser.UserStatus_UserStatusEnable
	//todo covert password
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return info, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return info, f.Error()
	}
	return info, nil
}

func (s *Server) UserUnregister(_ context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	data, err := codec.Encode(structs.UserUnregisterRequestType, info)
	if err != nil {
		return info, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return info, f.Error()
	}
	return info, nil
}

func (s *Server) UserLogin(ctx context.Context, req *pbuser.LoginRequest) (*pbuser.UserInfo, error) {
	info, err := s.UserLoad(ctx, req.Uid)
	if err != nil {
		return info, err
	}
	//todo covert password
	if req.Password != info.Password {
		return info, ErrorUserNotExistOrPasswordNotMatch
	}
	return info, nil
}

func (s *Server) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	info, err := s.UserLoad(ctx, req.Uid)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	info.Status = req.GetStatus()
	data, err := codec.Encode(structs.UserRegisterRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, err
}

func (s *Server) UserList(ctx context.Context, _ *emptypb.Empty) (*pbuser.UserListResponse, error) {
	resp := &pbuser.UserListResponse{}
	err := s.state.Scan(ctx, []byte(structs.UsersBucketName), func(k, v []byte) error {
		item := &pbuser.UserInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, err
}

func (s *Server) UserLoad(ctx context.Context, uid string) (*pbuser.UserInfo, error) {
	info := &pbuser.UserInfo{}
	data, err := s.state.Get(ctx, []byte(structs.UsersBucketName), []byte(uid))
	if err != nil {
		return info, err
	}
	err = codec.Decode(data, info)
	if err != nil {
		return info, err
	}
	if info.Uid == "" {
		return info, ErrorUserNotExist
	}
	return info, nil
}

func (s *Server) UserAddNamespace(_ context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.UserAddNamespaceRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UserDelNamespace(_ context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.UserDelNamespaceRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
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
	return resp, err
}
