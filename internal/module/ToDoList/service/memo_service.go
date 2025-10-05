package service

import (
	"SService/internal/module/ToDoList/dto"
	"SService/internal/module/ToDoList/model"
	dto2 "SService/internal/module/common/dto"
	common_model "SService/internal/module/common/model"
	"SService/pkg/database"
)

func CreateMemo(req dto.MemoCreateDTO) (*dto.MemoListItemDTO, error) {
	memo := &model.Memo{
		UserUUID:   req.UserUUID,
		Content:    req.Content,
		IsComplete: false,
		SortOrder:  req.SortOrder,
	}

	if err := database.DB.Create(memo).Error; err != nil {
		return nil, err
	}

	memoDTO := &dto.MemoListItemDTO{
		ID:         memo.ID,
		Content:    memo.Content,
		IsComplete: memo.IsComplete,
	}
	return memoDTO, nil // 返回完整的创建对象（包含 ID）
}

// 查找所有,GetMemoList
func GetMemoList(req dto.MemoListQueryDTO, userUUID common_model.UUID) (dto2.PaginationResponse, error) {
	// 默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	var total int64
	err := database.DB.Model(&model.Memo{}).
		Where("user_uuid = ?", userUUID).
		Count(&total).Error
	if err != nil {
		return dto2.PaginationResponse{}, err
	}
	var memoList []*model.Memo
	err = database.DB.Model(&model.Memo{}).
		Where("user_uuid = ?", userUUID).
		Order("sort_order ASC").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&memoList).Error
	if err != nil {
		return dto2.PaginationResponse{}, err
	}
	return dto2.PaginationResponse{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		Data:     memoList,
	}, nil
}
