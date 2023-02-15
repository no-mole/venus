package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/venus/api/kv"
	"github.com/no-mole/venus/agent/venus/server"
)

func Router(s server.Server) *gin.Engine {
	router := gin.New()

	pprof.Register(router)
	group := router.Group("/api/v1")

	kvGroup := group.Group("/kv")
	kvGroup.PUT("", kv.Put(s))
	return router
}
