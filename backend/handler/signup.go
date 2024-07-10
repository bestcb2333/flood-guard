package handler

import (
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var request struct {
		Username string
		Password string
		Email    string
		Authcode string
	}

	if err := c.BindJSON(&request); err != nil {
		util.Error(c, 400, "无法读取你的请求", err)
		return
	}

}
