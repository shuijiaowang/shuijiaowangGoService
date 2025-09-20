package service

import (
	"SService/internal/module/ToDoList/dto"
	"SService/internal/module/ToDoList/model"
	"SService/pkg/database"
)

func AddMemo(req dto.MemoCreateDTO) (*dto.MemoListItemDTO, error) {
	memo := &model.Memo{
		UserUUID:   req.UserUUID,
		Content:    req.Content,
		IsComplete: false,
	}

	if err := database.DB.Create(memo).Error; err != nil {
		return nil, err
	}

	memoDTO := &dto.MemoListItemDTO{
		ID:         memo.ID,
		Content:    memo.Content,
		IsComplete: memo.IsComplete,
		//出现错误
		//CreateTimeDTO: dto.CreateTimeDTO{
		//	CreatedAt: memo.CreatedAt,
		//},
		//CreatedAt: memo.CreatedAt,
	}
	memoDTO.CreatedAt = memo.CreatedAt // 直接赋值
	return memoDTO, nil                // 返回完整的创建对象（包含 ID）
}
