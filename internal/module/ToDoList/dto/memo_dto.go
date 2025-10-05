package dto

import (
	"SService/internal/module/common/dto"
)

/*
创建备忘录:
uuid用于后续传参补充结构体，无需传参
content内容是必传字段
sortOrder排序字段,这个字段理应是要传参的，但仍需要进行验证，如果是0，或是重复了，需要查最大的值+1000再存储？
*/
type MemoCreateDTO struct {
	dto.BaseReqDTO
	Content   string  `json:"content" binding:"required,min=1"` // 必传内容
	SortOrder float64 `json:"sortOrder" `
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
	dto.PaginationRequest
	IsComplete *bool `json:"is_complete"` // 可选筛选：只看已完成/未完成
}
type MemoListItemDTO struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
}
