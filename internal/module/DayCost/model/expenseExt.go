package model

import common_model "SService/internal/module/common/model"

// 这里不需要用户id，那这个的增删改查岂不是可以无限制？用户A发起向用户B修改的请求
type ExpenseExt struct {
	ExpenseID        int                    `gorm:"uniqueIndex;not null;comment:关联消费ID"` // 非指针（必填）
	ExpenseType      int8                   `gorm:"not null;comment:类型(0:时间型,1:数量型)"`    // 非指针（必填）
	StartDate        common_model.JSONDate  `gorm:"type:date;not null;comment:开始使用日期"`   // 非指针（必填）
	EstimatedEndDate *common_model.JSONDate `gorm:"comment:预计使用天数(时间型)"`                 // 指针（可选）
	EndDate          *common_model.JSONDate `gorm:"type:date;comment:实际结束日期"`            // 指针（可选）
	TotalQuantity    *float64               `gorm:"type:decimal(10,4);comment:总数量(数量型)"` // 指针（可选）
	Remaining        *float64               `gorm:"type:decimal(10,4);comment:剩余量(数量型)"` // 指针（可选）
	Status           int8                   `gorm:"default:1;comment:状态(0:进行中,1:已结束)"`   // 非指针（有默认值）
	common_model.BaseModel
}
