package service

import (
	"SService/internal/module/DayCost/repository"
)

type BaseService struct {
	baseRepo *repository.BaseRepository
}

func NewBaseService() *BaseService {
	return &BaseService{
		baseRepo: &repository.BaseRepository{},
	}
}

func (b *BaseService) CheckExpenseExtOwner(userID, expenseID int) error {
	return b.baseRepo.CheckExpenseOwner(expenseID, userID)
}
