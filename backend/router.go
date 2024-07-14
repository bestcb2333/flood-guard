package main

import (
	"github.com/bestcb2333/FloodGuard/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router() {

	r := gin.Default()

	r.GET("/captcha", handler.GetCaptcha)
	r.GET("/email", handler.GetMail)
	r.POST("/login", handler.AuthCaptcha, handler.Login)
	r.POST("/signup", handler.Signup)

	r.GET("/get/:path", handler.SelectRecord)
	r.POST("/edit/:path", handler.EditRecord)
	r.POST("/delete/:path", handler.DeleteRecord)

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
