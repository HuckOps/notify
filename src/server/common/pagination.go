package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pagination(ctx *gin.Context) (skip, limit int, err error) {
	skip, err1 := strconv.Atoi(ctx.DefaultQuery("skip", "0"))
	limit, err2 := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err1 != nil || err2 != nil {
		err = errors.New("Query error")
	}
	//findOptions.SetSkip(int64(skip))
	//findOptions.SetLimit(int64(limit))
	return
}
