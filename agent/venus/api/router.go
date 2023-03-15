package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/no-mole/venus/agent/docs"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/api/access_key"
	"github.com/no-mole/venus/agent/venus/api/kv"
	"github.com/no-mole/venus/agent/venus/api/namespace"
	"github.com/no-mole/venus/agent/venus/api/service"
	"github.com/no-mole/venus/agent/venus/api/user"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/metrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(s *venus.Server, a auth.Authenticator) *gin.Engine {
	//do not validate
	binding.Validator.Engine().(*validator.Validate).SetTagName("noBinding")

	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				output.Json(ctx, fmt.Errorf("%v", err), nil)
			}
		}()
		ctx.Next()
	})
	router.NoRoute(func(ctx *gin.Context) {
		output.Json(ctx, errors.New("no router"), nil)
		return
	})
	router.POST("/api/v1/login", Login(s))

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(MustLogin(s, a))
	group := router.Group("/api/v1")
	router.GET("/metrics", metrics.Collector.HttpHandler())
	router.Use(metrics.Collector.HttpRequestTotal(), metrics.Collector.HttpRequestDurationTime())

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
	userGroup.GET("", user.List(s))
	userGroup.GET("/:uid/namespace", user.NamespaceList(s))
	userGroup.POST("/:uid", user.Add(s))
	userGroup.PUT("/:uid", user.ChangePassword(s))

	accessKeyGroup := group.Group("/access_key")
	accessKeyGroup.GET("", access_key.List(s))
	accessKeyGroup.GET("/:ak/namespace", access_key.NamespaceList(s))
	accessKeyGroup.POST("/:namespace/:alias", access_key.Gen(s))
	accessKeyGroup.DELETE("/:ak", access_key.Del(s))
	accessKeyGroup.POST("/login/:ak", access_key.Login(s))
	accessKeyGroup.PUT("/:ak", access_key.ChangeStatus(s))

	authGroup := group.Group("/oauth2")
	authGroup.GET("/callback", Callback(s, a))
	return router
}
