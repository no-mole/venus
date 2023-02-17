package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/venus/api/kv"
	"github.com/no-mole/venus/agent/venus/prometheus"
	"github.com/no-mole/venus/agent/venus/server"
)

func Router(s server.Server) *gin.Engine {
	router := gin.New()

	router.Use(prometheus.NewPrometheusHandle(s.PrometheusServer()))

	group := router.Group("/api/v1")

	kvGroup := group.Group("/kv")
	kvGroup.PUT("", kv.Put(s))
	return router
}
