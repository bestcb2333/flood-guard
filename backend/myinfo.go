package main

import (
	"github.com/gin-gonic/gin"
)

func GetMyInfo(c *gin.Context, u *User, _ struct{}) {

	c.JSON(200, Resp("登录成功", nil, u))
}
