package repository

import (
	"SService/internal/module/DayCost/model"
	"SService/pkg/database"
)

// expense_ext_repo.go 继承基础Repository
type ExpenseExtRepository struct {
	BaseRepository // 嵌入基础Repository，复用CheckExpenseOwner
}

// h.expenseExtRepo.AddExpenseExt(req)
// 函数声明开始
func (r *ExpenseExtRepository) AddExpenseExt(expenseExt *model.ExpenseExt) error {
	tx := database.DB.Create(expenseExt)
	return tx.Error
}

// expenseExt,err:= h.expenseExtRepo.GetExpenseExtById(id)
func (r *ExpenseExtRepository) GetExpenseExtById(id int) (*model.ExpenseExt, error) {

	expenseExt := &model.ExpenseExt{}
	tx := database.DB.First(expenseExt, id)
	return expenseExt, tx.Error
}
