// 文件作用：HTTP 响应工具：提供统一成功/失败 JSON 结构封装。
package http

import "github.com/gin-gonic/gin"

// Success 返回统一的成功响应结构。
func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
		"data": data,
	})
}

// Failure 返回统一的失败响应结构。
func Failure(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"code": status,
		"msg":  message,
	})
}
