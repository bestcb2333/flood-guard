package main

import (
	"github.com/bestcb2333/FloodGuard/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router() {

	r := gin.Default()

	r.POST("/login", handler.Login)
	r.POST("/signup", handler.Signup)

	r.GET("/get/user", handler.GetUser)
	r.GET("/get/region", handler.GetRegion)
	r.GET("/get/floodevent", handler.GetFloodEvent)
	r.GET("/get/historydata", handler.GetHistoryData)
	r.GET("/get/notice", handler.GetNotice)

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
