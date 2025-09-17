// pkg/util/request.go
package util

import (
	"github.com/gin-gonic/gin"
)

// BindAndValidate 绑定并验证请求参数
func BindAndValidate(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		Result(c, 400, "无效的请求格式: "+err.Error(), nil)
		return false
	}
	return true
}
