package route

import (
	"github.com/HuckOps/notify/src/server/handler/middleware"
	"github.com/HuckOps/notify/src/server/handler/rbac"
	"github.com/gin-gonic/gin"
)

func Admin(e *gin.Engine) {

	rbacRouteMap := e.Group("/admin/rbac")
	{
		rbacRouteMap.Use(middleware.AuthMiddleware())
		rbacRouteMap.GET("/user", rbac.GetUserHandler)
		rbacRouteMap.POST("/user", rbac.PostUserHandler)
		rbacRouteMap.DELETE("/user", rbac.DeleteUserHandler)
		rbacRouteMap.GET("/active", rbac.GetActiveHandler)
		rbacRouteMap.POST("/active", rbac.PostActiveHandler)
	}
}
