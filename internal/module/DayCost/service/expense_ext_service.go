package service

import (
	"SService/internal/module/DayCost/model"
	"SService/internal/module/DayCost/repository"
)

type ExpenseExtService struct {
	expenseExtRepo *repository.ExpenseExtRepository
	expenseRepo    *repository.ExpenseRepository
}

func NewExpenseExtService() *ExpenseExtService {
	return &ExpenseExtService{
		expenseExtRepo: &repository.ExpenseExtRepository{},
	}
}

// err := h.expenseExtService.AddExpenseExt(userID,req)
func (h *ExpenseExtService) AddExpenseExt(userID int, req *model.ExpenseExt) error {
	err := h.expenseRepo.UpdateIsExtended(req.ExpenseID, userID, true)
	if err != nil {
		return err
	}
	err = h.expenseExtRepo.AddExpenseExt(req)
	if err != nil {
		return err
	}
	return nil
}
func (h *ExpenseExtService) GetExpenseExtById(id int) (*model.ExpenseExt, error) {
	expenseExt, err := h.expenseExtRepo.GetExpenseExtById(id)
	if err != nil {
		return nil, err // Repo层可能返回 gorm.ErrRecordNotFound
	}
	return expenseExt, nil

}
