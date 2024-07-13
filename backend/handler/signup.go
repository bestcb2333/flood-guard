package handler

import (
	"github.com/bestcb2333/FloodGuard/database"
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

	if !AuthMail(request.Authcode, request.Email) {
		util.Error(c, 400, "邮箱验证码不正确", nil)
		return
	}

	if err := DB.First(
		new(database.User), "username = ?", request.Username,
	); err == nil {
		util.Error(c, 400, "这个用户名已被注册", nil)
		return
	}

	if err := DB.Create(&database.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Admin:    false,
	}).Error; err != nil {
		util.Error(c, 400, "用户创建失败", err)
		return
	}

	util.Info(c, 200, "用户创建成功", nil)
}
