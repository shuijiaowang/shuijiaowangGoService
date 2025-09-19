package model

import (
	common_model "SService/internal/module/common/model"
	"time"
)

// HabitCheckin 习惯打卡模型
type HabitCheckin struct {
	UserUUID               common_model.UUID `gorm:"type:binary(16);not null;index:idx_user_uuid;comment:用户UUID"` // 新增
	HabitID                int               `gorm:"type:int unsigned;not null;index:idx_habit_id;comment:关联习惯ID"`
	CheckinTime            time.Time         `gorm:"type:datetime;not null;comment:打卡时间"`
	common_model.BaseModel                   // 嵌入基础模型

}

// TableName 指定表名
func (HabitCheckin) TableName() string {
	return "habit_checkins"
}
