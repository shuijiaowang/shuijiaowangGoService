package service

import (
	"SService/internal/module/DayCost/dto"
	"SService/internal/module/DayCost/model"
	"SService/internal/module/DayCost/repository"
)

type ExpenseService struct {
	expenseRepo *repository.ExpenseRepository
}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{
		expenseRepo: &repository.ExpenseRepository{},
	}
}

// 这里得一个一个参数传递，还是封装个vo对象
func (s *ExpenseService) AddExpense(expense *model.Expense) {
	s.expenseRepo.AddExpense(expense)
	return
}
func (s *ExpenseService) GetExpenseById(expenseID string, userID string) (*model.Expense, error) {
	// 直接调用Repo的方法，传入两个ID
	expense, err := s.expenseRepo.FindByIDAndUserID(expenseID, userID)
	if err != nil {
		return nil, err // Repo层可能返回 gorm.ErrRecordNotFound
	}
	return expense, nil
}

func (s *ExpenseService) ListExpense(userID string) ([]*model.Expense, error) { // 返回切片
	expenses, err := s.expenseRepo.ListExpense(userID)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

// 条件分页查询
func (s *ExpenseService) ListExpenseByCondition(query dto.ExpensePagesQuery) ([]dto.ExpenseDto, int64, error) {
	expenses, total, err := s.expenseRepo.ListByCondition(query)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.ExpenseDto
	for _, exp := range expenses {
		responses = append(responses, dto.ToResultExpense(exp))
	}

	return responses, total, nil
}

func (s *ExpenseService) UpdateExpense(expense *model.Expense) error {
	return s.expenseRepo.UpdateExpense(expense)
}

// 删
func (s *ExpenseService) DeleteExpense(expenseID string, userID string) error {
	return s.expenseRepo.DeleteExpense(expenseID, userID)
}

// 恢复
func (s *ExpenseService) RecoverExpense(expenseID string, userID string) error {
	return s.expenseRepo.RecoverExpense(expenseID, userID)
}

// Statistic 获取指定月份的每日收支统计
func (s *ExpenseService) Statistic(month string, userID string) ([]dto.ExpenseDay, error) {
	return s.expenseRepo.StatisticByMonth(month, userID)
}
