package model

import (
	common_model "SService/internal/module/common/model"
	"time"
)

// Task 任务模型(支持父子任务)
type Task struct {
	UserUUID               common_model.UUID `gorm:"type:binary(16);not null;index:idx_user_uuid;comment:用户UUID"`
	Title                  string            `gorm:"type:varchar(100);not null;comment:任务名称"`
	ParentTaskID           *int              `gorm:"index:idx_parent_task;comment:父任务ID(自关联)"` // 指针类型允许null
	IsCompleted            bool              `gorm:"type:tinyint(1);not null;default:0;comment:是否完成"`
	StartTime              *time.Time        `gorm:"type:datetime;comment:开始时间"`
	EndTime                *time.Time        `gorm:"type:datetime;comment:结束时间"`
	common_model.BaseModel                   // 嵌入基础模型
}

// TableName 指定表名
func (Task) TableName() string {
	return "task"
}
