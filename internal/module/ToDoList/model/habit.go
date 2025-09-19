package model

import (
	common_model "SService/internal/module/common/model"
	"time"
)

// Habit 习惯模型
type Habit struct {
	UserUUID               common_model.UUID `gorm:"type:binary(16);not null;index:idx_user_uuid;comment:用户UUID"`
	Title                  string            `gorm:"type:varchar(100);not null;comment:习惯名称"`
	IntervalValue          int               `gorm:"not null;comment:时间间隔(单位:天)"`
	StartDate              time.Time         `gorm:"type:date;not null;comment:开始日期"`
	EndDate                *time.Time        `gorm:"type:date;comment:结束日期"`
	IsActive               bool              `gorm:"type:tinyint(1);not null;default:1;comment:是否激活"`
	common_model.BaseModel                   // 嵌入基础模型

}

// TableName 指定表名
func (Habit) TableName() string {
	return "habit"
}
