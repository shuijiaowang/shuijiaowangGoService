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

// 在AuthHandler结构体所在的handler文件中添加

func (h *AuthHandler) Register(c *gin.Context) {
	// 定义注册请求参数结构体
	var req struct {
		Username string `json:"username" binding:"required,min=1,max=20"` // 用户名长度限制
		Password string `json:"password" binding:"required,min=1"`        // 密码长度限制
	}

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Result(c, 400, "无效的请求格式：用户名至少3位，密码至少6位", nil)
		return
	}

	// 调用服务层注册方法
	err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		// 根据错误信息返回对应提示（例如用户名已存在）
		util.Result(c, 400, err.Error(), nil)
		return
	}

	// 注册成功
	util.Result(c, 200, "注册成功，请登录", nil)
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
	token, err := util.GenerateToken(user.ID, user.Username, user.UserUUID)
	if err != nil {
		util.Result(c, 500, "生成令牌失败", nil)
		return
	}

	// 返回token和用户信息
	util.Result(c, 200, "登录成功", gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"user_uuid": user.UserUUID, // 可返回给前端用于展示或后续操作
		"token":     token,
	})
}

// test
func (h *AuthHandler) Test(c *gin.Context) {
	util.Result(c, 200, "ok", gin.H{})
}
