package handler

import (
	"SService/internal/module/user/service"
	"SService/pkg/util"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Result(c, 404, "无效的请求格式", nil)
		return
	}

	user, ok := h.authService.Login(req.Username, req.Password)
	if !ok {
		util.Result(c, 401, "用户名或密码错误", nil)
		return
	}

	// 生成JWT令牌
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		util.Result(c, 500, "生成令牌失败", nil)
		return
	}

	// 返回token和用户信息
	util.Result(c, 200, "登录成功", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"token":    token,
	})
}

// test
func (h *AuthHandler) Test(c *gin.Context) {
	util.Result(c, 200, "ok", gin.H{})
}
