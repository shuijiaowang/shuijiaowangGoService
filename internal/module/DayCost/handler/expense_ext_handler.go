package handler

import (
	"SService/internal/module/DayCost/model"
	"SService/internal/module/DayCost/repository"
	"SService/internal/module/DayCost/service"
	"SService/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseExtHandler struct {
	BaseHandler                                  // 继承BaseHandler
	expenseExtService *service.ExpenseExtService // 内部创建要比传参方便一些
}

// 构造
func NewExpenseExtHandler() *ExpenseExtHandler {
	return &ExpenseExtHandler{
		BaseHandler:       *NewBaseHandler(), // 关键：初始化父结构体
		expenseExtService: service.NewExpenseExtService(),
	}
}
func (h *ExpenseExtHandler) AddExpenseExt(c *gin.Context) {
	userID := h.GetUserID(c) //获取用户id
	var req repository.ExpenseExtDto
	h.Bind(c, &req)
	expenseExt := &model.ExpenseExt{
		ExpenseID:        req.ExpenseID,
		ExpenseType:      req.ExpenseType,
		StartDate:        req.StartDate,
		EstimatedEndDate: req.EstimatedEndDate,
		EndDate:          req.EndDate,
		TotalQuantity:    req.TotalQuantity,
		Remaining:        req.Remaining,
	}
	// 检查是否为该用户
	h.CheckExpenseExtOwner(c, userID, req.ExpenseID)
	// 添加
	err := h.expenseExtService.AddExpenseExt(userID, expenseExt)
	if err != nil {
		util.Result(c, 500, "添加失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "添加成功", nil)
}

// GetExpenseExtById
func (h *ExpenseExtHandler) GetExpenseExtById(c *gin.Context) {
	userID := h.GetUserID(c)

	idStr := c.Param("id") //查询id
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.Result(c, 400, "id参数错误", nil)
		return
	}
	// 检查是否为该用户
	h.CheckExpenseExtOwner(c, userID, id)

	expenseExt, err := h.expenseExtService.GetExpenseExtById(id)
	if err != nil {
		util.Result(c, 500, "获取失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "获取成功", expenseExt)

}
