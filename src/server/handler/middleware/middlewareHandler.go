package middleware

import (
	"context"
	"fmt"
	"github.com/HuckOps/notify/pkg/restful"
	"github.com/HuckOps/notify/src/db/mongo"
	"github.com/HuckOps/notify/src/model/mysql"
	"go.mongodb.org/mongo-driver/bson"

	//"github.com/dgrijalva/jwt-go"
	"github.com/HuckOps/notify/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if user, err := jwt.ParseToken(token); err != nil {
			restful.FailedResponse(ctx, 400, restful.NoPermission, map[string]interface{}{}, "token error")
			ctx.Abort()
		} else {
			ctx.Set("user", user.User)
			ctx.Next()
		}
	}
}

func PermissionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenant := ctx.GetHeader("TENANT")
		//user, _ := ctx.Get("user")
		//userInstance := user.(jwt.User)
		tenantInstance := mysql.Tenant{}
		mongo.Mongo.DB().Collection("tenant").FindOne(context.TODO(), bson.M{"code": tenant}).Decode(&tenantInstance)
		result := map[string]interface{}{}
		r, err := mongo.Mongo.DB().Collection("user_to_sub").Aggregate(context.TODO(), []bson.M{
			bson.M{
				"$match": []bson.D{},
			},
			bson.M{
				"$lookup": bson.M{
					"from":         "sub",
					"localField":   "sub_id",
					"foreignField": "_id",
					"as":           "source",
				},
			},
		})
		fmt.Println(err)
		if r.Next(context.Background()) {
			err := r.Decode(&result)
			fmt.Println(result, err)
		}
	}
}
