package main

import (
	"time"

	"github.com/bestcb2333/FloodGuard/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router() {

	r := gin.Default()

	//允许CORS跨域
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/captcha", handler.GetCaptcha)
	r.GET("/email", handler.GetMail)
	r.GET("/get/:path", handler.SelectRecord)
	r.GET("gpt", handler.Gpt)

	r.Use(handler.PostMidWare)
	r.POST("/login", handler.AuthCaptcha, handler.Login)
	r.POST("/signup", handler.Signup)
	r.POST("/edit/:path", handler.EditRecord)
	r.POST("/delete/:path", handler.DeleteRecord)
	r.POST("/gpt", handler.Gpt)

	if viper.GetString("SSL_ENABLE") == "true" {
		r.RunTLS(
			":"+viper.GetString("port"),
			viper.GetString("SSL_CERTIFICATE"),
			viper.GetString("SSL_KEY"),
		)
	} else {
		r.Run(":" + viper.GetString("port"))
	}
}
