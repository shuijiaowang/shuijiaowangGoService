package dto

import (
	common_model "SService/internal/module/common/model"
	"time"
)

// BaseReqDTO 公共的请求参数结构体,这里暂时只包含用户的uuid，这个好吗？
type BaseReqDTO struct {
	UserUUID common_model.UUID `json:"user_uuid"` // 公共的用户标识
}
type CreateTimeDTO struct {
	CreatedAt time.Time `json:"created_at"` // 创建时间
}
