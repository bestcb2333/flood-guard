package handler

import (
	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := util.ParseJSON(c, &request); err != nil {
		util.Error(c, 400, "用户请求读取失败", err)
		return
	}

	if request.Username == "" || request.Password == "" {
		util.Error(c, 400, "请提供完整的用户名和密码", nil)
		return
	}

	var user db.User
	if err := DB.First(
		&user, "username = ?", request.Username,
	).Error; err != nil {
		util.HandleQueryErr(c, "你尚未注册", nil)
		return
	}

	if user.Password != request.Password {
		util.Error(c, 400, "密码不正确", nil)
		return
	}

	token, err := util.GetJwt(user.ID)
	if err != nil {
		util.Error(c, 500, "服务器错误", err)
		return
	}

	util.Info(c, 200, "登录成功", gin.H{
		"token":    token,
		"username": user.Username,
		"admin":    user.Admin,
		"email":    user.Email,
	})
}
