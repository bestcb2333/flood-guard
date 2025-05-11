package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginHandlerDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddLoginRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.POST("/login", Preload(
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{JSON: true},
		},
		new(LoginHandlerDTO),
		func(c *gin.Context, _ *User, r *LoginHandlerDTO) {

			if r.Username == "" || r.Password == "" {
				c.JSON(400, Resp("用户名和密码不完整", nil, nil))
				return
			}

			var user User
			if err := pbc.DB.First(
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

			token, err := GetJwt(user.ID, pbc.JWTKey)
			if err != nil {
				c.JSON(500, Resp("服务器错误", err, nil))
				return
			}

			c.JSON(200, Resp("登录成功", nil, token))
		},
	))
}
