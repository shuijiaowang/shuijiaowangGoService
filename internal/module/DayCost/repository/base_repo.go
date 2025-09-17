// repository/base_repo.go（基础Repository）
package repository

import (
	"SService/internal/module/DayCost/model"
	"SService/pkg/database"
	"fmt"
)

type BaseRepository struct{}

// CheckExpenseOwner 校验ExpenseID是否属于当前用户（通用方法）
func (r *BaseRepository) CheckExpenseOwner(expenseID, userID int) error {
	var count int64
	if err := database.DB.Model(&model.Expense{}).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", expenseID, userID).
		Count(&count).Error; err != nil {
		return fmt.Errorf("权限校验失败: %w", err)
	}
	return nil
}
