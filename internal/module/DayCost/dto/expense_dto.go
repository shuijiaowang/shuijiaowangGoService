// SService/internal/dto/expense_dto.go
package dto

import (
	"SService/internal/module/DayCost/model"
	common "SService/internal/module/common/model"
)

// ExpenseDto （不含UserIdcreateAt、UpdatedAt、DeletedAt）
type ExpenseDto struct {
	ID              int             `json:"id"`
	UserID          int             `json:"user_id"`
	Note            string          `json:"note" binding:"required,max=100"`                          // 物品名称/摘要（必填，最长100字符）
	Amount          float64         `json:"amount" binding:"required"`                                // 金额（必填，大于0）
	Remarks         string          `json:"remarks" binding:"omitempty,max=500"`                      // 备注（可选，最长500字符）
	ExpenseDate     common.JSONDate `json:"expense_date" binding:"required" time_format:"2006-01-02"` // 消费日期（必填）
	Category        int8            `json:"category" binding:"required,gte=0,lte=9"`                  // 分类（必填，0-9范围内）
	IsExtended      bool            `json:"is_extended"`
	TransactionType int8            `json:"transaction_type" binding:"required,gte=0,lte=1`
}

//这俩就没必要了
// CreateExpenseRequest 接收前端新增消费的请求参数
// returnExpenseResponse 返回给前端的数据（不含userId）

// ExpenseQuery 查询消费的请求参数 //form:"" 表示使用form表单提交？
type ExpenseQuery struct {
	UserID          int             `json:"user_id"`          // 用户ID（内部用，前端不传递）
	NoteLike        string          `json:"note_like"`        // 物品名称模糊查询
	RemarksLike     string          `json:"remarks_like"`     // 备注模糊查询
	MinAmount       *float64        `json:"min_amount"`       // 最小金额（>=）
	MaxAmount       *float64        `json:"max_amount"`       // 最大金额（<=）
	StartDate       common.JSONDate `json:"start_date"`       // 开始日期（>=）
	EndDate         common.JSONDate `json:"end_date"`         // 结束日期（<=）
	Category        int8            `json:"category"`         // 分类（>0有效）
	IsExtended      *bool           `json:"is_extended"`      // 是否扩展消费（指针非nil时作为条件）
	SortBy          string          `json:"sort_by"`          // 排序字段（如expense_date、amount）
	SortOrder       string          `json:"sort_order"`       // 排序方向（asc/desc，默认asc）
	IsDeleted       bool            `json:"is_deleted"`       // 查看回收站
	TransactionType *int8           `json:"transaction_type"` //交易类型（0:支出,1:收入）
	//ExpenseExtType      *int8          `json:"expense_ext_type"`      // 扩展消费类型（0:时间型,1:数量型，仅IsExtended=true有效）
	//ExpenseExtStatus    *int8          `json:"expense_ext_status"`    // 扩展消费状态（0:进行中,1:已结束，仅IsExtended=true有效）
}

// 消费列表查询请求（包含条件+分页）
type ExpensePagesQuery struct {
	ExpenseQuery      // 嵌入业务查询条件
	PaginationRequest // 嵌入分页参数
}

type ExpenseDay struct {
	Date    string  `json:"date"`    // 日期，格式：yyyy-mm-dd
	Expense float64 `json:"expense"` // 当日总支出
	Income  float64 `json:"income"`  // 当日总收入
}
type ExpenseMonth struct {
	Month        string       `json:"month"`         // 月份，格式：yyyy-mm
	TotalExpense float64      `json:"total_expense"` // 当月总支出
	TotalIncome  float64      `json:"total_income"`  // 当月总收入
	DailyDetails []ExpenseDay `json:"daily_details"` // 当月每日明细
}

// 传入expense,返回espenseDto,除去(userid/createAt/updatedAt/deletedAt)
func ToResultExpense(expense *model.Expense) ExpenseDto {
	return ExpenseDto{
		ID:              expense.ID,
		Note:            expense.Note,
		Amount:          expense.Amount,
		Remarks:         expense.Remarks,
		ExpenseDate:     expense.ExpenseDate,
		Category:        expense.Category,
		IsExtended:      expense.IsExtended,
		TransactionType: expense.TransactionType,
	}
}
