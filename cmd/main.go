package main

import (
	"SService/pkg/corn"
	"SService/pkg/database"
	"SService/routes"
	"log"
)

func main() {
	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 初始化定时任务
	c, err := cron.Init()
	if err != nil {
		log.Fatalf("定时任务初始化失败: %v", err)
	}
	defer c.Stop() // 程序退出时停止

	// 创建路由
	r := routes.SetupRouter()

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
