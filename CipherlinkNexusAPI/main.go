package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 版本信息端点
	r.GET("/version/core", func(c *gin.Context) {
		// 获取version.json的绝对路径
		versionFile := filepath.Join("versions", "core.json")

		// 读取文件内容
		data, err := os.ReadFile(versionFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "无法读取版本信息文件",
			})
			return
		}

		// 验证JSON格式
		var result interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "版本信息文件格式错误",
			})
			return
		}

		// 设置内容类型并返回JSON
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(data))
	})

	// 启动服务器
	log.Println("API服务器已启动，监听端口 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
