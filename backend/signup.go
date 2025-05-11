package main

import (
	"github.com/gin-gonic/gin"
)

type SignupDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddSignupRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.POST("/signup", Preload(
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
		},
		new(SignupDTO),
		func(c *gin.Context, _ *User, r *SignupDTO) {

			email := c.GetHeader("X-Email-Number")
			if err := pbc.DB.First(
				new(User), "email = ?", email,
			).Error; err == nil {
				c.JSON(400, Resp("这个邮箱已经被使用", nil, nil))
				return
			}

			if err := pbc.DB.First(
				new(User), "name = ?", r.Username,
			); err == nil {
				c.JSON(400, Resp("这个用户名已被使用", nil, nil))
				return
			}

			if err := pbc.DB.Create(&User{
				Name:     r.Username,
				Password: r.Password,
				Email:    email,
				Admin:    false,
			}).Error; err != nil {
				c.JSON(500, Resp("用户创建失败", err, nil))
				return
			}

			c.JSON(200, Resp("用户创建成功", nil, nil))
		},
	))
}
