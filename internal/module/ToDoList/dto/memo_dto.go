package dto

import (
	"SService/internal/module/common/dto"
)

type MemoCreateDTO struct {
	dto.BaseReqDTO
	Content string `json:"content" binding:"required,min=1"` // 必传内容
}

// 2. 更新备忘录：只包含可修改的字段（用指针区分「未传值」和「传空值」）
type MemoUpdateDTO struct {
	dto.BaseReqDTO
	ID         uint    `json:"id" binding:"required"` // 必传ID（定位要更新的记录）
	Content    *string `json:"content"`               // 可选：传则更新，不传则不修改
	IsComplete *bool   `json:"is_complete"`           // 可选：同上
}

// 3. 备忘录列表查询：可能需要分页+筛选，返回精简字段
type MemoListQueryDTO struct {
	dto.BaseReqDTO       // 嵌入用户标识（筛选当前用户的备忘录）
	IsComplete     *bool `json:"is_complete"`          // 可选筛选：只看已完成/未完成
	Page           int   `json:"page" binding:"min=1"` // 分页参数（复用公共分页逻辑）
	Size           int   `json:"size" binding:"min=1,max=100"`
}
type MemoListItemDTO struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
	dto.CreateTimeDTO
}
