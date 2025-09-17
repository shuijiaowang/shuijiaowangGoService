package handler

import (
	"SService/internal/module/DayCost/service"
	"SService/pkg/util"

	"github.com/gin-gonic/gin"
)

// BaseHandler 基础Handler，封装公共逻辑
type BaseHandler struct {
	baseService *service.BaseService
}

// 初始化
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		baseService: service.NewBaseService(),
	}
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

// 判断expenseID和userId是否匹配来判断是否有权限？
func (h *BaseHandler) CheckExpenseExtOwner(c *gin.Context, userID int, expenseID int) {
	err := h.baseService.CheckExpenseExtOwner(userID, expenseID)
	if err != nil {
		// 权限校验失败时，抛出自定义AppError，由全局异常中间件统一响应
		panic(util.NewAppError(401, "无权限操作该消费记录", err))
	}
}
