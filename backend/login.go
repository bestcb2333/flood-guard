package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context, _ *User, r struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) {

	if r.Username == "" || r.Password == "" {
		c.JSON(400, Resp("用户名和密码不完整", nil, nil))
		return
	}

	var user User
	if err := DB.First(
		&user, "name = ?", r.Username,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(400, Resp("你尚未注册", err, nil))
		return
	} else if err != nil {
		c.JSON(500, Resp("用户查询失败", err, nil))
		return
	}

	if user.Password != r.Password {
		c.JSON(400, Resp("密码不正确", nil, nil))
		return
	}

	token, err := GetJwt(user.ID)
	if err != nil {
		c.JSON(500, Resp("服务器错误", err, nil))
		return
	}

	c.JSON(200, Resp("登录成功", nil, token))
}
