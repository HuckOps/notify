package rbac

import (
	"github.com/HuckOps/notify/pkg/restful"
	"github.com/HuckOps/notify/src/db/mysql"
	mysqlModel "github.com/HuckOps/notify/src/model/mysql"
	"github.com/HuckOps/notify/src/server/common"
	"github.com/gin-gonic/gin"
)

func GetPermission(ctx *gin.Context) {
	tenant := ctx.GetHeader("TENANT")
	result := make([]mysqlModel.CasbinRule, 0)
	user, _ := ctx.Get("user")
	groups_code := []string{}
	for _, g := range user.(mysqlModel.User).Groups {
		if g.Tenant.Code == tenant {
			groups_code = append(groups_code, g.Code)
		}
	}
	mysql.MySQL.DB().Model(&result).Where("v0 in ? and v3 = ?", groups_code, tenant).Scan(&result)

	restful.SuccessResponse(ctx, result, "")
}

func GetUserHandler(ctx *gin.Context) {
	skip, limit, err := common.Pagination(ctx)
	var search string
	search = ctx.DefaultQuery("search", "")
	if err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	res := restful.MakeItemResponse[User]()
	fetch := mysql.MySQL.DB().Model(&User{})
	if search != "" {
		fetch = fetch.Where("username = ? or name = ? or telephone = ?", search, search, search)
	}
	mysql.SearchByPagination[User](fetch, skip, limit, &res)
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
	user.Model = userModel.Model
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
	active := restful.MakeItemResponse[mysqlModel.Active]()
	mysql.SearchByPagination[mysqlModel.Active](mysql.MySQL.DB().Model(&mysqlModel.Tenant{}), skip, limit, &active)
	restful.SuccessResponse(ctx, active, "")
}

func PostActiveHandler(ctx *gin.Context) {
	act := mysqlModel.Active{}
	if err := ctx.ShouldBindJSON(&act); err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError,
			map[interface{}]interface{}{}, "request error")
		return
	}
	mysql.MySQL.DB().Create(&act)
	restful.SuccessResponse(ctx, act, "")
}

func GetTenant(ctx *gin.Context) {
	skip, limit, err := common.Pagination(ctx)
	if err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	tenant := restful.MakeItemResponse[mysqlModel.Tenant]()
	mysql.SearchByPagination[mysqlModel.Tenant](mysql.MySQL.DB().Model(&mysqlModel.Tenant{}), skip, limit, &tenant)
	responseItem := []Tenant{}
	for _, item := range tenant.Items {
		responseItem = append(responseItem, Tenant{item, len(item.Users), 0})
	}
	restful.SuccessResponse(ctx, restful.ItemResponse[Tenant]{Total: tenant.Total, Items: responseItem}, "")
}

func PostTenant(ctx *gin.Context) {
	tenant := mysqlModel.Tenant{}
	if err := ctx.ShouldBindJSON(&tenant); err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, err.Error())
		return
	}
	mysql.MySQL.DB().Create(&tenant)
	restful.SuccessResponse(ctx, tenant, "")
}

func DelTenant(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError, map[string]interface{}{}, "id未获取")
		return
	}
	err := mysql.MySQL.DB().Where("id = ?", id).Delete(&mysqlModel.Tenant{}).Error
	if err != nil {
		return
	}
	restful.SuccessResponse(ctx, "delete success", "")
}
