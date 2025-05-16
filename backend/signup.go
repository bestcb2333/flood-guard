package main

import (
	"time"

	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
)

type SignupDTO struct {
	Email    string `json:"email"`
	Authcode string `json:"authcode"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddSignupRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.POST("/signup", p.Preload(
		&p.Config[SignupDTO]{
			Base: bc,
			Bind: &p.BindConfig{JSON: true},
		},
		func(c *gin.Context, _ *User, r *SignupDTO) {

			if err := bc.DB.First(
				new(User), "email = ?", r.Email,
			).Error; err == nil {
				c.JSON(400, Resp("这个邮箱已经被使用", nil, nil))
				return
			}

			EmailCodeData.Mu.Lock()
			value, exist := EmailCodeData.Data[r.Email]
			EmailCodeData.Mu.Unlock()
			if !exist {
				c.JSON(500, bc.Resp("未申请验证码或已过期", nil, nil))
				return
			}

			if time.Now().After(value.Expiration) {
				c.JSON(400, bc.Resp("验证码已过期", nil, nil))
				return
			}

			if r.Authcode != value.Code {
				c.JSON(400, bc.Resp("验证码不正确", nil, nil))
			}

			if err := bc.DB.First(
				new(User), "name = ?", r.Username,
			); err == nil {
				c.JSON(400, Resp("这个用户名已被使用", nil, nil))
				return
			}

			if err := bc.DB.Create(&User{
				Name:     r.Username,
				Password: r.Password,
				Email:    r.Email,
				Admin:    false,
			}).Error; err != nil {
				c.JSON(500, Resp("用户创建失败", err, nil))
				return
			}

			c.JSON(200, Resp("用户创建成功", nil, nil))
		},
	))
}
