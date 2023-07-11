package route

import (
	"github.com/HuckOps/notify/src/server/handler/middleware"
	"github.com/HuckOps/notify/src/server/handler/rbac"
	"github.com/gin-gonic/gin"
)

func Admin(e *gin.Engine) {
	permissionRouteMap := e.Group("/")
	{
		permissionRouteMap.Use(middleware.AuthMiddleware())
		permissionRouteMap.GET("/permission", rbac.GetPermission)
	}

	rbacRouteMap := e.Group("/admin/rbac")
	{

		rbacRouteMap.Use(middleware.AuthMiddleware())
		rbacRouteMap.Use(middleware.PermissionMiddleware())
		rbacRouteMap.GET("/user", rbac.GetUserHandler)
		rbacRouteMap.POST("/user", rbac.PostUserHandler)
		rbacRouteMap.DELETE("/user", rbac.DeleteUserHandler)
		rbacRouteMap.GET("/active", rbac.GetActiveHandler)
		rbacRouteMap.POST("/active", rbac.PostActiveHandler)

	}
	tenantRouteMap := e.Group("/admin")
	{
		tenantRouteMap.GET("/tenant", rbac.GetTenant)
		tenantRouteMap.POST("/tenant", rbac.PostTenant)
		tenantRouteMap.DELETE("/tenant", rbac.DelTenant)
	}
}
