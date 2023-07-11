package route

import (
	"github.com/HuckOps/notify/src/server/handler/middleware"
	"github.com/gin-gonic/gin"
)

func Project(e *gin.Engine) {
	projectPermission := e.Group("/project")
	{
		projectPermission.Use(middleware.AuthMiddleware())
		projectPermission.Use(middleware.PermissionMiddleware())
		projectPermission.GET("/users")
	}
}
