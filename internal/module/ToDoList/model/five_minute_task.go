package model

import (
	common_model "SService/internal/module/common/model"
)

// FiveMinuteTask 五分钟任务模型
type FiveMinuteTask struct {
	UserUUID               common_model.UUID `gorm:"type:binary(16);not null;index:idx_user_uuid;comment:用户UUID"` // 新增
	TaskID                 int               `gorm:"type:int unsigned;not null;index:idx_task_id;comment:关联任务ID"`
	Description            string            `gorm:"type:text;not null;comment:任务描述"`
	IsCompleted            bool              `gorm:"type:tinyint(1);not null;default:0;comment:是否完成"`
	common_model.BaseModel                   // 嵌入基础模
}

// TableName 指定表名
func (FiveMinuteTask) TableName() string {
	return "five_minute_task"
}
