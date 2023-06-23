package restful

import "github.com/gin-gonic/gin"

type ResponseCode int

const (
	Success ResponseCode = iota
	NoPermission
	RequestError
	ServerError
)

const (
	OK               = 200
	BadRequest       = 400
	PermissionDenied = 403
	NotFound         = 404
)

type ItemResponse[T interface{}] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}

type Response[Data interface{}] struct {
	Code ResponseCode `json:"code" example:0`
	MSG  string       `json:"msg" example:""`
	Data Data         `json:"data"`
}

func SuccessResponse(ctx *gin.Context, v interface{}, msg string) {
	ctx.JSON(OK, Response[interface{}]{Code: Success, MSG: msg, Data: v})
	ctx.Abort()
}

func FailedResponse(ctx *gin.Context, statusCode int, responseCode ResponseCode, v interface{}, msg string) {
	ctx.JSON(statusCode, Response[interface{}]{Code: responseCode, MSG: msg, Data: v})
	ctx.Abort()
}

func MakeItemResponse[T interface{}]() ItemResponse[T] {
	item := ItemResponse[T]{}
	item.Items = make([]T, 0)
	return item
}
