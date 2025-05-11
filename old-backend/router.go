package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", Ping)
	r.GET("/captcha", GetCaptcha)
	r.POST("/login", LoginHandler)
	r.POST("/signup", AuthCaptcha, AuthEmail, SignupHandler)
	r.GET("/mail", Preload(SendMail, PARAM))
	r.GET("/myinfo", Preload(GetMyInfo, LOGIN, "Region"))
	r.GET("/factors", GetRiskFactor)
	r.GET("/regions", GetRegions)

	r.GET("/notices", Preload(GetNotices, PARAM))

	r.GET("/histories", Preload(GetHistoryData, PARAM))
	r.GET("/trends", GetHistoryTrend)

	r.GET("/sensors", Preload(GetSensors, PARAM))
	r.GET("/resources", GetResources)
	r.GET("/events", Preload(GetEvents, PARAM))
	r.GET("/users", Preload(GetUsers, PARAM))
	r.GET("/deepseek", Preload(GetDeepSeek, LOGIN, "Analyses"))
	r.POST("/deepseek", Preload(PostDeepSeek, LOGIN|JSON, "Analyses"))
	r.POST("/upload/region", Preload(UploadRegion, LOGIN|ADMIN|BIND))
	r.POST("/upload/event", Preload(UploadEvent, LOGIN|ADMIN|BIND))

	return r
}
