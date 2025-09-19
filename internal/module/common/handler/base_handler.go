// SService/internal/module/common/handler/base_handler.go
package handler

import (
	common "SService/internal/module/common/model"
	"SService/pkg/util"

	"github.com/gin-gonic/gin"
)

// BaseHandler 基础Handler，封装公共逻辑，Bind方法/GetUserID方法
type BaseHandler struct {
}

// 初始化
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// Bind 通用参数绑定方法，绑定失败时抛出自定义错误（由全局中间件处理）
func (h *BaseHandler) Bind(c *gin.Context, req interface{}) {
	if err := c.ShouldBindJSON(req); err != nil {
		panic(util.NewAppError(400, "无效的请求格式: "+err.Error(), err))
	}
	// 绑定成功时无返回值，直接继续执行
}

// GetUserID 从上下文获取用户ID，封装重复的断言和错误处理
func (h *BaseHandler) GetUserID(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		panic(util.NewAppError(401, "未获取到用户信息", nil))
	}

	id, ok := userID.(int) //断言用户ID类型
	if !ok {
		panic(util.NewAppError(401, "无效的用户ID类型", nil))
	}
	return id
}
func (h *BaseHandler) GetUserUUID(c *gin.Context) common.UUID {
	userID, exists := c.Get("userUUID")
	if !exists {
		panic(util.NewAppError(401, "未获取到用户信息", nil))
	}

	id, ok := userID.(common.UUID) //断言用户ID类型
	if !ok {
		panic(util.NewAppError(401, "无效的用户UUID类型", nil))
	}
	return id
}
