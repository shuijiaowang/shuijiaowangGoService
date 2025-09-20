package dto

import (
	common "SService/internal/module/common/model"
)

// 联合DTO：同时包含普通消费更新和拓展消费信息

type ExpenseExtDto struct {
	ID               int              `json:"id" comment:"扩展ID"`
	ExpenseID        int              `json:"expense_id" comment:"关联消费ID"`
	ExpenseType      int8             `json:"expense_type" comment:"类型(0:时间型,1:数量型)"`
	StartDate        common.JSONDate  `json:"start_date" comment:"开始使用日期"`
	EstimatedEndDate *common.JSONDate `json:"estimated_end_date" comment:"预计结束日期(时间型)"`
	EndDate          *common.JSONDate `json:"end_date" comment:"实际结束日期"`
	TotalQuantity    *float64         `json:"total_quantity" comment:"总数量(数量型)"`
	Remaining        *float64         `json:"remaining" comment:"剩余量(数量型)"`
	Status           int8             `json:"status" comment:"状态(0:进行中,1:已结束)"`
}
