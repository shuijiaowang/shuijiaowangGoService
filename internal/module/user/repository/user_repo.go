package repository

import (
	"SService/internal/module/user/model"
	"SService/pkg/database"
)

// 用户仓库结构体
type UserRepository struct{}

// 根据用户名查询用户
func (r *UserRepository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	// Where 条件查询 + First 获取第一条记录
	result := database.DB.Where("username = ?", name).First(&user)
	return &user, result.Error
}

// 创建用户
func (r *UserRepository) CreateUser(user *model.User) error {
	// Create 插入记录
	result := database.DB.Create(user)
	return result.Error
}
