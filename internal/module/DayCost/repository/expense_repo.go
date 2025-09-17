package repository

import (
	"SService/internal/module/DayCost/dto"
	"SService/internal/module/DayCost/model"
	"SService/pkg/database"
	"fmt"
	"time"
)

type ExpenseRepository struct{}

func (r *ExpenseRepository) AddExpense(expense *model.Expense) {
	database.DB.Create(expense) //GORM 在插入记录后，可能需要修改传入的结构体实例，以便将数据库生成的值（如自增ID、时间戳等）回填到结构体中。所以需要传入指针
	return
}

func (r *ExpenseRepository) FindByIDAndUserID(expenseID string, userID string) (*model.Expense, error) {
	var expense model.Expense
	// 在查询条件中同时包含主键ID和用户ID！
	result := database.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense)

	if result.Error != nil {
		return nil, result.Error // 如果没找到，这里会返回 gorm.ErrRecordNotFound
	}
	return &expense, nil
}

func (r *ExpenseRepository) ListExpense(userID string) ([]*model.Expense, error) {
	var expenses []*model.Expense
	result := database.DB.Where("user_id = ?", userID).Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

// 条件分页查询
func (r *ExpenseRepository) ListByCondition(query dto.ExpensePagesQuery) ([]*model.Expense, int64, error) {
	var expenses []*model.Expense //返回切片
	db := database.DB.Model(&model.Expense{}).Where("user_id = ?", query.UserID)

	// 处理软删除筛选：只看删除的 / 只看未删除的 /查看回收站
	if query.IsDeleted {
		// 只看已删除：需要 Unscoped() 取消默认过滤，再筛选 deleted_at 不为空
		db = db.Unscoped().Where("deleted_at IS NOT NULL")
	} else {
		// 只看未删除：默认筛选 deleted_at 为空（GORM 软删除默认行为，显式写更清晰）
		db = db.Where("deleted_at IS NULL") //虽重复也会被优化掉//这里更好可读
	}
	// 构建查询条件

	// 构建查询条件
	// 模糊查询（支持空字符串不生效）
	if query.NoteLike != "" {
		db = db.Where("note LIKE ?", "%"+query.NoteLike+"%")
	}
	if query.RemarksLike != "" {
		db = db.Where("remarks LIKE ?", "%"+query.RemarksLike+"%")
	}

	// 价格查询（支持0值，指针非nil表示用户传递了参数，包括0）
	if query.MinAmount != nil { // 先判断指针是否非空（用户传递了该参数）
		db = db.Where("amount >= ?", *query.MinAmount) // 解引用获取值
	}
	if query.MaxAmount != nil { // 同理处理MaxAmount
		db = db.Where("amount <= ?", *query.MaxAmount)
	}

	//日期查询（支持固定日期：StartDate和EndDate设为同一个值即可）
	if !query.StartDate.ToTime().IsZero() {
		db = db.Where("expense_date >= ?", query.StartDate)
	}
	if !query.EndDate.ToTime().IsZero() {
		db = db.Where("expense_date <= ?", query.EndDate)
	}

	// 分类查询
	if query.Category > 0 { // 原逻辑是>0，
		db = db.Where("category = ?", query.Category)
	}

	// 是否为扩展记录（指针非nil时生效，支持false值）
	if query.IsExtended != nil {
		db = db.Where("is_extended = ?", query.IsExtended)
	}
	// 交易类型（支出还是收入）
	if query.TransactionType != nil {
		db = db.Where("transaction_type = ?", query.TransactionType)
	}

	// 获取总记录数（必须先Count再排序/分页，否则总条数会受分页影响）
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询总记录数失败: %w", err)
	}

	// 处理排序（支持单字段排序，多字段可后续扩展）
	if query.SortBy != "" {
		sortOrder := query.SortOrder
		if sortOrder == "" || (sortOrder != "asc" && sortOrder != "desc") {
			sortOrder = "asc"
		}
		db = db.Order(query.SortBy + " " + sortOrder)
	} else {
		// 默认按消费日期倒序、ID倒序（最新的记录在前）
		db = db.Order("expense_date desc, id desc")
	}

	// 分页处理（计算偏移量，注意Page可能为0的边界情况）
	offset := (query.Page - 1) * query.PageSize
	if offset < 0 {
		offset = 0 // 避免Page=0时出现负偏移
	}

	// 执行查询并检查错误
	if err := db.Limit(query.PageSize).Offset(offset).Find(&expenses).Error; err != nil {
		return nil, 0, fmt.Errorf("查询消费记录失败: %w", err)
	}
	return expenses, total, nil
}

// 修改
func (r *ExpenseRepository) UpdateExpense(expense *model.Expense) error {
	result := database.DB.Model(&model.Expense{}).
		Select("Note", "Amount", "Remarks", "ExpenseDate", "Category", "IsExtended").
		Where("id = ? AND user_id = ?", expense.ID, expense.UserID).
		Updates(expense)

	if result.Error != nil {
		return fmt.Errorf("更新消费记录失败: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("记录不存在或无权修改")
	}

	return nil
}

// 更新IsExtended
func (r *ExpenseRepository) UpdateIsExtended(id int, userID int, isExtended bool) error {
	return database.DB.Model(&model.Expense{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_extended", isExtended).Error
}

// 删除
func (r *ExpenseRepository) DeleteExpense(expenseID string, userID string) error {
	result := database.DB.Where("id = ? AND user_id = ?", expenseID, userID).Delete(&model.Expense{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 恢复
func (r *ExpenseRepository) RecoverExpense(expenseID string, userID string) error {
	return database.DB.Unscoped().Model(&model.Expense{}).
		Where("id = ? AND user_id = ?", expenseID, userID).
		Update("deleted_at", nil).Error
}

// StatisticByMonth 按月份统计每日收支
func (r *ExpenseRepository) StatisticByMonth(month string, userID string) ([]dto.ExpenseDay, error) {
	// 计算月份的起止日期
	startDate := month + "-01"
	// 解析月份最后一天
	lastDay, err := getLastDayOfMonth(month)
	if err != nil {
		return nil, err
	}
	endDate := month + "-" + lastDay

	// SQL查询：按天分组统计收支
	var stats []struct {
		Date    string  `gorm:"column:date"`
		Expense float64 `gorm:"column:expense"`
		Income  float64 `gorm:"column:income"`
	}

	err = database.DB.Raw(`
		SELECT 
			DATE_FORMAT(expense_date, '%Y-%m-%d') as date,
			SUM(CASE WHEN transaction_type = 0 THEN amount ELSE 0 END) as expense,
			SUM(CASE WHEN transaction_type = 1 THEN amount ELSE 0 END) as income
		FROM expense
		WHERE user_id = ? 
			AND expense_date BETWEEN ? AND ?
			AND deleted_at IS NULL
		GROUP BY expense_date
		ORDER BY date
	`, userID, startDate, endDate).Scan(&stats).Error

	if err != nil {
		return nil, fmt.Errorf("统计查询失败: %w", err)
	}

	// 转换为DTO并补全当月所有日期（包括无数据的日期）
	statMap := make(map[string]dto.ExpenseDay)
	for _, s := range stats {
		statMap[s.Date] = dto.ExpenseDay{
			Date:    s.Date,
			Expense: s.Expense,
			Income:  s.Income,
		}
	}

	// 生成当月所有日期并填充数据
	var result []dto.ExpenseDay
	current, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	for !current.After(end) {
		dateStr := current.Format("2006-01-02")
		if stat, ok := statMap[dateStr]; ok {
			result = append(result, stat)
		} else {
			result = append(result, dto.ExpenseDay{
				Date:    dateStr,
				Expense: 0,
				Income:  0,
			})
		}
		current = current.AddDate(0, 0, 1)
	}

	return result, nil
}

// 辅助函数：获取月份最后一天
func getLastDayOfMonth(month string) (string, error) {
	t, err := time.Parse("2006-01", month)
	if err != nil {
		return "", err
	}
	// 下个月第一天减一天就是当月最后一天
	lastDay := t.AddDate(0, 1, -1).Format("02")
	return lastDay, nil
}
