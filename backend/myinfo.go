package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
)

func RegGetMyinfoHandler(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/myinfo", p.Preload(
		&p.Config[struct{}]{
			Base:       bc,
			Permission: p.Login,
		},
		func(c *gin.Context, u *User, r *struct{}) {

			c.JSON(200, bc.Resp("获取用户信息成功", nil, u))
		},
	))
}
