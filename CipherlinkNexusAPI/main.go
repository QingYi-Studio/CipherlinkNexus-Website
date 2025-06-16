package main

import (
	"CipherlinkNexusAPI/routes"
	"log"
)

func main() {
	// 初始化路由
	r := routes.SetupRouter()

	// 启动服务器
	log.Println("API服务器已启动，监听端口 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
