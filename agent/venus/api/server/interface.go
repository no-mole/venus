package server

import (
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"github.com/no-mole/venus/proto/pbuser"
)

type Server interface {
	server.Server
	GetSysConfig() *pbsysconfig.SysConfig
	UserSync(info *pbuser.UserInfo) (*pbuser.LoginResponse, error)
}
