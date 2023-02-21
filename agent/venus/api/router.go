package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/no-mole/venus/agent/docs"
	"github.com/no-mole/venus/agent/venus/api/access_key"
	"github.com/no-mole/venus/agent/venus/api/kv"
	"github.com/no-mole/venus/agent/venus/api/namespace"
	"github.com/no-mole/venus/agent/venus/api/service"
	"github.com/no-mole/venus/agent/venus/api/user"
	"github.com/no-mole/venus/agent/venus/server"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(s server.Server) *gin.Engine {
	//docs.SwaggerInfo.Host = xxxx//todo
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	namespaceGroup.POST("/:namespace/access_key/:ak", namespace.AccessKeyAdd(s))
	namespaceGroup.DELETE("/:namespace/access_key/:ak", namespace.AccessKeyDel(s))
	namespaceGroup.GET("/:namespace/access_key", namespace.AccessKeyList(s))

	serviceGroup := group.Group("/service")
	serviceGroup.GET("/:namespace", service.List(s))
	serviceGroup.GET("/:namespace/:service_name", service.Versions(s))
	serviceGroup.GET("/:namespace/:service_name/:service_version", service.Endpoints(s))

	userGroup := group.Group("/user")
	userGroup.POST("/:uid", user.Add(s))
	userGroup.PUT("/:uid", user.ChangePassword(s))
	userGroup.POST("/login/:uid", user.Login(s))

	accessKeyGroup := group.Group("/access_key")
	accessKeyGroup.POST("/:ak", access_key.Gen(s))
	accessKeyGroup.DELETE("/:ak", access_key.Del(s))
	accessKeyGroup.POST("/login/:ak", access_key.Login(s))
	accessKeyGroup.GET("", access_key.List(s))
	accessKeyGroup.PUT("/:ak", access_key.ChangeStatus(s))
	return router
}
