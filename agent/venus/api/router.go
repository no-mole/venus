package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/venus/api/kv"
	"github.com/no-mole/venus/agent/venus/api/namespace"
	"github.com/no-mole/venus/agent/venus/api/service"
	"github.com/no-mole/venus/agent/venus/api/user"
	"github.com/no-mole/venus/agent/venus/server"
)

func Router(s server.Server) *gin.Engine {
	router := gin.New()

	group := router.Group("/api/v1")

	kvGroup := group.Group("/kv")
	kvGroup.PUT("/:namespace/:key", kv.Put(s))
	kvGroup.GET("/:namespace", kv.List(s))
	kvGroup.DELETE("/:namespace/:key", kv.Del(s))
	kvGroup.GET("/:namespace/:key", kv.Fetch(s))

	namespaceGroup := group.Group("/namespace")
	namespaceGroup.POST("/:namespace", namespace.Add(s))
	namespaceGroup.DELETE("/:namespace", namespace.Del(s))
	namespaceGroup.GET("", namespace.List(s))
	namespaceGroup.POST("/:namespace/user/:uid", namespace.UserAdd(s))
	namespaceGroup.DELETE("/:namespace/user/:uid", namespace.UserDel(s))
	namespaceGroup.GET("/:namespace/user", namespace.UserList(s))

	serviceGroup := group.Group("/service")
	serviceGroup.GET("/:namespace", service.List(s))
	serviceGroup.GET("/:namespace/:service_name", service.Versions(s))
	serviceGroup.GET("/:namespace/:service_name/:service_version", service.Endpoints(s))

	userGroup := group.Group("/user")
	userGroup.POST("/:uid", user.Add(s))
	userGroup.PUT("/:uid", user.ChangePassword(s))
	userGroup.POST("/login/:uid", user.Login(s))
	return router
}
