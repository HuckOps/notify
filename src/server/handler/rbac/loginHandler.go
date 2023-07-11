package rbac

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"github.com/HuckOps/notify/pkg/jwt"
	"github.com/HuckOps/notify/pkg/restful"
	"github.com/HuckOps/notify/src/db/mysql"
	mysqlModel "github.com/HuckOps/notify/src/model/mysql"
	_ "github.com/HuckOps/notify/src/server/docs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login 登录处理事务
// @Summary 创建镜像迁移任务
// @Description 创建镜像迁移任务
// @Tags Users
// @Accept json
// @Produce json
// @Success 200  {object}  LoginResponse "请求成功"
// @Router /login [post]
func Login(ctx *gin.Context) {
	req := LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restful.FailedResponse(ctx, restful.BadRequest, restful.RequestError,
			map[interface{}]interface{}{}, "request error")
		return
	}
	// 密码反序列化
	decodeString, _ := base64.StdEncoding.DecodeString(req.Password)
	privateKey, _ := ctx.Get("priKey")
	result, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), decodeString)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	password := string(result)
	user := mysqlModel.User{}
	err = mysql.MySQL.DB().Model(&user).Where("username = ?", req.UserName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		restful.FailedResponse(ctx, 401, restful.NoPermission, map[string]interface{}{}, "Not found user")
		return
	}
	if user.Password != password {
		restful.FailedResponse(ctx, 401, restful.RequestError,
			map[string]interface{}{}, "Password error")
		return
	}
	token, err := jwt.GetToken(jwt.User{ID: user.ID, UserName: user.UserName, Name: user.Name, Telephone: user.Telephone})
	if err != nil {
		restful.FailedResponse(ctx, 500, restful.ServerError, map[string]interface{}{}, err.Error())
		return
	}
	restful.SuccessResponse(ctx, gin.H{"token": token}, "")
}

func WhoamI(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if exist {
		restful.SuccessResponse(ctx, user.(mysqlModel.User), "")
	}
}
