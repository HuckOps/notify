package route

import (
	"github.com/HuckOps/notify/src/server/handler/middleware"
	"github.com/HuckOps/notify/src/server/handler/rbac"
	"github.com/gin-gonic/gin"
)

func User(e *gin.Engine) {
	loginRouteMap := e.Group("")
	{
		loginRouteMap.POST("login", rbac.Login)
		loginRouteMap.GET("whoami", middleware.AuthMiddleware(), rbac.WhoamI)
	}
	tenantRouteMap := e.Group("/tenant")
	{
		tenantRouteMap.Use(middleware.AuthMiddleware())
		tenantRouteMap.Use(middleware.PermissionMiddleware())
		tenantRouteMap.GET("")
	}
}
