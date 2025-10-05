package model

import (
	model2 "SService/internal/module/common/model"
)

type User struct {
	ID       int         `gorm:"primaryKey;autoIncrement;comment:用户ID"`
	Username string      `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名"`
	Password string      `gorm:"type:varchar(100);not null;comment:加密密码"`
	UserUUID model2.UUID `json:"userUUID" gorm:"type:binary(16);uniqueIndex;not null;comment:用户唯一标识UUID"`
	model2.BaseModel
}
