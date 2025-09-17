package routes

import (
	handler3 "SService/internal/module/DayCost/handler"
	handler2 "SService/internal/module/user/handler"
	"SService/pkg/middleware" // 导入中间件包

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 跨域中间件（放在最前面）
	r.Use(middleware.Cors())
	// 注册全局异常处理中间件
	r.Use(middleware.ErrorHandler())

	authHandler := handler2.NewAuthHandler()
	expenseHandler := handler3.NewExpenseHandler()
	expenseExtHandler := handler3.NewExpenseExtHandler()
	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", authHandler.Login)
		//userGroup.POST("/register", authHandler.Register)
		userGroup.GET("/test", authHandler.Test)
	}

	// 需要认证的路由组
	// 增删改查
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTInterceptor()) // 应用JWT拦截器
	{
		// 消费记录路由（需要认证）
		expenseGroup := apiGroup.Group("/expenses")
		{
			expenseGroup.POST("/", expenseHandler.AddExpense)               // 添加消费记录
			expenseGroup.GET("/:id", expenseHandler.GetExpenseById)         // 获取单个消费记录
			expenseGroup.GET("/", expenseHandler.ListExpense)               // 获取消费记录列表
			expenseGroup.GET("/by", expenseHandler.ListExpenseByCondition)  // 获取消费记录列表
			expenseGroup.PUT("/", expenseHandler.UpdateExpense)             // 更新消费记录
			expenseGroup.DELETE("/:id", expenseHandler.DeleteExpense)       // 删除消费记录
			expenseGroup.PUT("/recover/:id", expenseHandler.RecoverExpense) //恢复
			expenseGroup.GET("/statistic", expenseHandler.Statistic)        //默认统计
			expenseGroup.GET("/test", expenseHandler.Test)                  //默认统计
		}
		// 消费拓展路由（需要认证）
		expenseExtGroup := apiGroup.Group("/expense-ext") //路由使用驼峰？
		{
			expenseExtGroup.POST("/", expenseExtHandler.AddExpenseExt)       // 添加
			expenseExtGroup.GET("/:id", expenseExtHandler.GetExpenseExtById) // 获取单个
			//expenseExtGroup.GET("/", expenseExtHandler.ListExpenseExt)              // 获取列表
			//expenseExtGroup.GET("/by", expenseExtHandler.ListExpenseExtByCondition) // 获取列表条件分页
			//expenseExtGroup.PUT("/", expenseExtHandler.UpdateExpenseExt)            // 更新
			//expenseExtGroup.DELETE("/:id", expenseExtHandler.DeleteExpenseExt)      // 删除
			//expenseExtGroup.PUT("/recover/:id", expenseExtHandler.RecoverExpenseExt)
			//expenseExtGroup.GET("/statistic", expenseExtHandler.Statistic)
		}
	}

	return r
}
