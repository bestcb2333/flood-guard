package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 输出错误信息到客户端和控制台和数据库
func Error(c *gin.Context, status int, msg string, err error) {
	c.AbortWithStatusJSON(status, RespBody(msg, nil))
	if err != nil {
		msg += err.Error()
		fmt.Println("出现了错误: " + msg)
	}
}

// 输出成功信息
func Info(c *gin.Context, status int, msg string, data any) {
	c.AbortWithStatusJSON(status, RespBody(msg, data))
}

// 制作响应体
func RespBody(msg string, data any) gin.H {
	return gin.H{"msg": msg, "data": data}
}

// 处理查询数据库时发生的错误
func HandleQueryErr(c *gin.Context, msg string, err error) {
	if err == gorm.ErrRecordNotFound {
		Error(c, 200, msg, nil)
	} else {
		Error(c, 500, "出现了错误", err)
	}
}
