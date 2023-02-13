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
	kvGroup.PUT("", kv.Put(s))
	kvGroup.GET("/:namespace", kv.List(s))

	namespaceGroup := group.Group("/namespace")
	namespaceGroup.POST("", namespace.Add(s))
	namespaceGroup.DELETE("/:namespace", namespace.Del(s))
	namespaceGroup.GET("", namespace.List(s))
	namespaceGroup.POST("/user", namespace.UserAdd(s))
	namespaceGroup.DELETE("/user", namespace.UserDel(s))
	namespaceGroup.GET("/user/:namespace", namespace.UserList(s))

	serviceGroup := group.Group("/service")
	serviceGroup.GET("/:namespace", service.List(s))
	serviceGroup.GET("/:namespace/:service_name", service.Versions(s))

	userGroup := group.Group("/user")
	userGroup.POST("", user.Add(s))
	userGroup.PUT("", user.ChangePassword(s))

	return router
}
