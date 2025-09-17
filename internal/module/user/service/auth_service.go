package service

import (
	"SService/internal/module/user/model"
	"SService/internal/module/user/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: &repository.UserRepository{},
	}
}

func (s *AuthService) Register(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(name, password string) (*model.User, bool) {
	user, err := s.userRepo.FindUserByName(name)
	if err != nil {
		return nil, false
	}

	// 简化密码验证（实际应使用bcrypt）
	if user.Password != password {
		return nil, false
	}

	return user, true
}
