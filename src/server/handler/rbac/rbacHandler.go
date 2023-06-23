package rbac

import (
	"context"
	"github.com/HuckOps/notify/pkg/restful"
	"github.com/HuckOps/notify/src/db/mongo"
	"github.com/HuckOps/notify/src/db/mysql"
	mysqlModel "github.com/HuckOps/notify/src/model/mysql"
	"github.com/HuckOps/notify/src/server/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserHandler(ctx *gin.Context) {
	skip, limit, err := common.Pagination(ctx)
	if err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	res := restful.MakeItemResponse[User]()
	mysql.SearchByPagination[User](mysql.MySQL.DB().Model(&User{}), skip, limit, &res)
	restful.SuccessResponse(ctx, res, "")
}

func PostUserHandler(ctx *gin.Context) {
	user := User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError,
			map[interface{}]interface{}{}, "request error")
		return
	}
	//
	userModel := mysqlModel.User{UserName: user.UserName, Name: user.Name, Telephone: user.Telephone, Password: "123456"}
	//_, err := mongo.Mongo.DB().Collection("user").InsertOne(context.TODO(), userModel)
	err := mysql.MySQL.DB().Model(&userModel).Create(&userModel).Error
	if err != nil {
		restful.FailedResponse(ctx, 500, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	user.ID = userModel.ID
	restful.SuccessResponse(ctx, user, "")

}

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, "id未获取")
		return
	}
	err := mysql.MySQL.DB().Where("id = ?", id).Delete(&mysqlModel.User{}).Error
	if err != nil {
		return
	}
	restful.SuccessResponse(ctx, "delete success", "")
}

func GetActiveHandler(ctx *gin.Context) {
	skip, limit, err := common.Pagination(ctx)
	if err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	active, _ := mongo.SearchByPagination(mongo.Mongo.DB().Collection("active"), skip, limit, bson.M{})
	res := restful.MakeItemResponse[Active]()
	if active.Next(context.Background()) {
		err = active.Decode(&res)
		if err != nil {
			restful.FailedResponse(ctx, 500, restful.RequestError, map[string]interface{}{}, err.Error())
		}
	}
	restful.SuccessResponse(ctx, res, "")
}

func PostActiveHandler(ctx *gin.Context) {
	act := Active{}
	if err := ctx.ShouldBindJSON(&act); err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError,
			map[interface{}]interface{}{}, "request error")
		return
	}
	_, err := mongo.Mongo.DB().Collection("active").InsertOne(context.TODO(), act)
	if err != nil {
		restful.FailedResponse(ctx, 500, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	restful.SuccessResponse(ctx, act, "")
}
