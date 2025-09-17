package middleware

import (
	"SService/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTInterceptor 验证JWT令牌的中间件
func JWTInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.Result(c, http.StatusUnauthorized, "未提供token", nil) //写响应，前端收到
			c.Abort()                                                //终止后续处理
			return                                                   //返回
		}

		// 检查格式是否为Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.Result(c, http.StatusUnauthorized, "token格式错误", nil)
			c.Abort()
			return
		}

		// 解析token
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			util.Result(c, http.StatusUnauthorized, "无效的token", nil)
			c.Abort()
			return
		}

		// 将用户ID存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
