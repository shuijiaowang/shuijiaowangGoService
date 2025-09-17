package model

import common_model "SService/internal/module/common/model"

type Expense struct {
	ID              int                   `gorm:"primaryKey;autoIncrement;comment:消费ID"`
	UserID          int                   `gorm:"not null;index;comment:关联用户ID"`
	Note            string                `gorm:"type:varchar(100);not null;comment:物品名称/消费摘要"`
	Amount          float64               `gorm:"type:decimal(10,2);not null;comment:消费金额"`
	Remarks         string                `gorm:"type:text;comment:详细备注(支持扩展标签)"`
	ExpenseDate     common_model.JSONDate `gorm:"type:date;not null;comment:消费日期"`
	Category        int8                  `gorm:"not null;comment:消费分类(0:餐饮,1:日用,2:交通...)"`
	IsExtended      bool                  `gorm:"default:false;comment:是否为扩展消费"`
	TransactionType int8                  `gorm:"not null;default:0;comment:交易类型(0:支出,1:收入,)"`
	common_model.BaseModel
}
