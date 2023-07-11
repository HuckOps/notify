package middleware

import (
	"github.com/HuckOps/notify/pkg/rbac"
	"github.com/HuckOps/notify/pkg/restful"
	mysqlEngine "github.com/HuckOps/notify/src/db/mysql"
	"github.com/HuckOps/notify/src/model/mysql"
	"reflect"
	"strings"

	//"github.com/dgrijalva/jwt-go"
	"github.com/HuckOps/notify/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 登录鉴权，获取用户的信息以及所属的租户
		token := ctx.GetHeader("TOKEN")
		if user, err := jwt.ParseToken(token); err != nil {
			restful.FailedResponse(ctx, 400, restful.NoPermission, map[string]interface{}{}, "token error")
			ctx.Abort()
		} else {
			u := mysql.User{}
			if err := mysqlEngine.MySQL.DB().Model(&mysql.User{}).Where("username = ?", user.User.UserName).Preload("Tenants").Preload("Groups").Preload("Groups.Tenant").First(&u).Error; err != nil {
				restful.FailedResponse(ctx, 400, restful.NoPermission, map[string]interface{}{}, err.Error())
				ctx.Abort()
				return
			}
			ctx.Set("user", u)
			ctx.Next()
		}
	}
}

func PermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取租户下的权限
		tenant := ctx.GetHeader("TENANT")
		user, _ := ctx.Get("user")

		// 超级管理员不做鉴权，直接pass
		if user.(mysql.User).IsAdmin {
			ctx.Next()
			return
		}

		tenantInstance := mysql.Tenant{}
		// 寻找当前匹配的租户
		for _, t := range user.(mysql.User).Tenants {
			if tenant == t.Code {
				tenantInstance = t
			}
		}

		if (reflect.DeepEqual(tenantInstance, mysql.Tenant{}) && !user.(mysql.User).IsAdmin) {
			restful.FailedResponse(ctx, 403, restful.NoPermission, map[string]interface{}{}, "dont have tenant permission")
			ctx.Abort()
			return
		}
		for _, group := range user.(mysql.User).Groups {
			if group.TenantID == tenantInstance.ID {
				if access, err := rbac.Enforce.Enforce(group.Code, ctx.Request.Method, strings.Split(ctx.Request.RequestURI, "?")[0], tenant); access && err == nil {
					ctx.Next()
					return
				}
			}
		}
		restful.FailedResponse(ctx, 403, restful.NoPermission, map[string]interface{}{}, "dont have operate permission")
		ctx.Abort()
	}
}
