package route

import (
	"github.com/HuckOps/notify/src/server/handler/middleware"
	"github.com/HuckOps/notify/src/server/handler/rbac"
	"github.com/gin-gonic/gin"
)

func User(e *gin.Engine) {
	loginRouteMap := e.Group("/login")
	{
		loginRouteMap.POST("", rbac.Login)
	}
	tenantRouteMap := e.Group("/tenant")
	{
		tenantRouteMap.Use(middleware.AuthMiddleware())
		tenantRouteMap.Use(middleware.PermissionMiddleware())
		tenantRouteMap.GET("")
	}
}
