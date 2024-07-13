package main

import (
	"github.com/bestcb2333/FloodGuard/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router() {

	r := gin.Default()

	r.GET("/captcha", handler.GetCaptcha)
	r.POST("/login", handler.AuthCaptcha, handler.Login)
	r.POST("/signup", handler.Signup)

	r.GET("/get/*any", handler.SelectRecord)
	r.POST("/edit/*any", handler.EditRecord)
	r.POST("/delete/*any", handler.DeleteRecord)

	if viper.GetBool("useSSL") {
		r.RunTLS(
			":"+viper.GetString("port"),
			viper.GetString("ssl.cert"),
			viper.GetString("ssl.key"),
		)
	} else {
		r.Run(":" + viper.GetString("port"))
	}
}
