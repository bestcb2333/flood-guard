package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 输出错误信息到客户端和控制台和数据库
func Error(c *gin.Context, status int, msg string, err error) {
	c.AbortWithStatusJSON(status, RespBody(msg, nil))
	if err != nil {
		msg += err.Error()
		fmt.Println("出现了错误: " + msg)
	}
}

// 制作响应体
func RespBody(msg string, data any) gin.H {
	return gin.H{"msg": msg, "data": data}
}
