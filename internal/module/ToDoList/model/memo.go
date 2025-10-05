package model

import (
	common_model "SService/internal/module/common/model"
)

// Memo 备忘录模型
type Memo struct {
	UserUUID               common_model.UUID `gorm:"type:binary(16);not null;index:idx_user_uuid;comment:用户UUID"`
	Content                string            `gorm:"type:text;not null;comment:备忘录内容"`
	IsComplete             bool              `gorm:"type:tinyint(1);not null;default:0;comment:是否完成(0:未完成,1:已完成)"`
	SortOrder              float64           `gorm:"type:float;not null;default:0;comment:排序字段"`
	common_model.BaseModel                   // 嵌入基础模型(ID、创建时间、更新时间、删除时间)
}

// TableName 指定表名
func (Memo) TableName() string {
	return "memo"
}
