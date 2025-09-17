package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS      = 200 // 成功
	BadRequest   = 400 //错误请求，参数错误
	Unauthorized = 401 //未授权，登陆失败
	NotFound     = 404 //资源不存在
	ServerError  = 500 //服务器内部错误
)

// 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 基础响应方法
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// 成功响应
func Success(c *gin.Context) {
	Result(c, SUCCESS, "ok", nil)
}

func SuccessWithMessage(c *gin.Context, message string) {
	Result(c, SUCCESS, message, nil)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "ok", data)
}

func SuccessDetailed(c *gin.Context, message string, data interface{}) {
	Result(c, SUCCESS, message, data)
}
