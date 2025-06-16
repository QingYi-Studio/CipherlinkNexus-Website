package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// GetVersionInfo 处理获取版本信息的请求
func GetVersionInfo(c *gin.Context) {
	// 获取version.json的绝对路径
	versionFile := filepath.Join("versions", "core.json")

	// 读取文件内容
	data, err := ioutil.ReadFile(versionFile)
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
}
