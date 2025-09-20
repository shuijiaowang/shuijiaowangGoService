package handler

import (
	dto2 "SService/internal/module/DayCost/dto"
	"SService/internal/module/DayCost/model"
	"SService/internal/module/DayCost/service"
	"SService/internal/module/common/dto"
	handler2 "SService/internal/module/common/handler"
	"SService/pkg/util"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	handler2.BaseHandler                         // 继承BaseHandler
	expenseService       *service.ExpenseService // 内部创建要比传参方便一些
}

// 构造
func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: service.NewExpenseService(),
	}
}

// 增
func (h *ExpenseHandler) AddExpense(c *gin.Context) {

	println("tianjia")
	// 复用BaseHandler的GetUserID，减少重复代码
	userID := h.GetUserID(c)

	// 2. 绑定并验证前端请求参数
	// 复用BaseHandler的Bind方法
	var req dto2.ExpenseDto
	h.Bind(c, &req)
	// 3. 转换DTO为数据库模型（只传递需要的字段）
	expense := &model.Expense{
		UserID:      userID, // 从上下文获取，前端无法篡改
		Note:        req.Note,
		Amount:      req.Amount,
		Remarks:     req.Remarks,
		ExpenseDate: req.ExpenseDate,
		Category:    req.Category,
		// IsExtended默认false，CreatedAt/UpdatedAt由数据库自动生成，无需手动设置
	}
	// 4. 调用服务层保存数据
	h.expenseService.AddExpense(expense)
	//if err != nil {
	//	util.Result(c, 500, "添加失败: "+err.Error(), nil)
	//	return
	//}
	util.Result(c, 200, "添加成功", nil)

}

// 查id
func (h *ExpenseHandler) GetExpenseById(c *gin.Context) {

	userID := h.GetUserID(c)
	//获取路径参数
	id := c.Param("id")
	// 将当前用户ID和要查询的ID一起传给Service
	expense, err := h.expenseService.GetExpenseById(id, strconv.Itoa(userID))
	if err != nil {
		util.Result(c, 403, err.Error(), nil) // 或者404，根据错误类型判断
		return
	}
	result := dto2.ToResultExpense(expense)
	util.Result(c, 200, "查询成功", result)
}

// 查所有
func (h *ExpenseHandler) ListExpense(c *gin.Context) {
	userID := h.GetUserID(c)
	expense, err := h.expenseService.ListExpense(strconv.Itoa(userID)) //返回切片
	if err != nil {
		util.Result(c, 403, "false", nil)
	}
	var expenseList []dto2.ExpenseDto
	for i := 0; i < len(expense); i++ {
		expenseList = append(expenseList, dto2.ToResultExpense(expense[i]))
	}
	util.Result(c, 200, "查询成功", expenseList)
}

// 条件查询+分页查询
func (h *ExpenseHandler) ListExpenseByCondition(c *gin.Context) {
	userID := h.GetUserID(c)
	var req dto2.ExpensePagesQuery
	h.Bind(c, &req)
	req.UserID = userID
	fmt.Println(req)
	expenses, total, err := h.expenseService.ListExpenseByCondition(req)
	if err != nil {
		util.Result(c, 500, "查询失败: "+err.Error(), nil)
		return
	}
	// 包装分页响应
	resp := dto.PaginationResponse{
		Total:    total,        //总页数
		Page:     req.Page,     //查询页码，第几页
		PageSize: req.PageSize, //每页条数
		Data:     expenses,     // 转换为DTO后的数据
	}
	util.Result(c, 200, "查询成功", resp)
}

// UpdateExpense
// 前端先获取，再修改保存
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	userID := h.GetUserID(c)
	var req dto2.ExpenseDto
	h.Bind(c, &req)

	req.UserID = userID
	expense := &model.Expense{
		ID:          req.ID,
		UserID:      userID,
		Note:        req.Note,
		Amount:      req.Amount,
		Remarks:     req.Remarks,
		ExpenseDate: req.ExpenseDate,
		Category:    req.Category,
		IsExtended:  req.IsExtended,
	}
	err := h.expenseService.UpdateExpense(expense)
	if err != nil {
		util.Result(c, 500, "更新失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "更新成功", nil)
}

// 删除
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userID := h.GetUserID(c)
	id := c.Param("id")
	err := h.expenseService.DeleteExpense(id, strconv.Itoa(userID))
	if err != nil {
		util.Result(c, 500, "删除失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "删除成功", nil)
}

// 恢复
func (h *ExpenseHandler) RecoverExpense(c *gin.Context) {
	userID := h.GetUserID(c)
	id := c.Param("id")
	err := h.expenseService.RecoverExpense(id, strconv.Itoa(userID))
	if err != nil {
		util.Result(c, 500, "恢复失败: "+err.Error(), nil)
		return
	}
	util.Result(c, 200, "恢复成功", nil)
}

//统计
//Statistic
//返回一个切片，默认返回该月的每一天的总支出，总收入
//传参的话应该传入月份？

func (h *ExpenseHandler) Statistic(c *gin.Context) {
	userID := h.GetUserID(c)
	// 1. 获取月份参数（格式：yyyy-mm），默认当前月份
	month := c.Query("month") // 建议用query参数更灵活，如/statistic?month=2024-05
	if month == "" {
		// 生成当前月份的 yyyy-mm 格式
		month = time.Now().Format("2006-01")
	}
	// 2. 验证月份格式是否正确
	_, err := time.Parse("2006-01", month)
	if err != nil {
		util.Result(c, 400, "月份格式错误，请使用yyyy-mm格式", nil)
		return
	}
	// 3. 调用服务层获取统计数据
	statisticData, err := h.expenseService.Statistic(month, strconv.Itoa(userID))
	if err != nil {
		util.Result(c, 500, "统计失败: "+err.Error(), nil)
		return
	}

	// 4. 返回统计结果
	util.Result(c, 200, "统计成功", statisticData)

}
func (h *ExpenseHandler) Test(c *gin.Context) {
	userID := h.GetUserID(c)
	util.Result(c, 200, "ok", userID)

}
