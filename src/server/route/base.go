package route

import (
	"crypto/rsa"
	_ "github.com/HuckOps/notify/src/server/docs"
	"github.com/gin-gonic/gin"
)

func Handler(e *gin.Engine) {
	User(e)
	Admin(e)
}

func SetPrivateKeyToContext(e *gin.Engine, pri *rsa.PrivateKey) {
	setContext := func(ctx *gin.Context) {
		ctx.Set("priKey", pri)
	}
	e.Use(setContext)

}
