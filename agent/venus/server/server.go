package server

import (
	"github.com/no-mole/venus/agent/venus/prometheus"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbservice"
	"github.com/no-mole/venus/proto/pbuser"
)

type Server interface {
	pbkv.KVServer
	pblease.LeaseServiceServer
	pbnamespace.NamespaceServiceServer
	pbuser.UserServiceServer
	pbservice.ServiceServer
	pbraftadmin.RaftAdminServer
	PrometheusServer() *prometheus.Prometheus
}
