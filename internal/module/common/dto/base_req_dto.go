package dto

import (
	common_model "SService/internal/module/common/model"
	"time"
)

type BaseReqDTO struct {
	UserUUID common_model.UUID `json:"user_uuid"` // 公共的用户标识
}
type CreateTimeDTO struct {
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
