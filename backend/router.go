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
	r.GET("/metadata", GetMetadata)
	r.GET("/captcha", GetCaptcha)
	r.GET("/map", Preload(GetMap, PARAM))
	r.POST("/login", Preload(Login, JSON))
	r.POST("/signup", AuthCaptcha, AuthEmail, Preload(Signup, JSON))
	r.GET("/mail", Preload(SendMail, PARAM))
	r.GET("/myinfo", Preload(GetMyInfo, LOGIN, "Region"))
	r.GET("/factor", GetRiskFactor)
	r.GET("/deepseek", Preload(GetDeepSeek, LOGIN, "Analyses"))
	r.POST("/deepseek", Preload(PostDeepSeek, LOGIN|JSON, "Analyses"))
	r.POST("/upload/region", Preload(UploadRegion, LOGIN|ADMIN|BIND))
	r.POST("/upload/event", Preload(UploadEvent, LOGIN|ADMIN|BIND))

	r.GET("/user", Get[User])
	r.GET("/region", Get[Region])
	r.GET("/event", Get[Event])
	r.GET("/history", Get[History])
	r.GET("/resource", Get[Resource])
	r.GET("/notice", Get[Notice])
	r.GET("/sensor", Get[Sensor])
	r.GET("/analysis", Get[Analysis])

	r.POST("/user", Preload(Set[User], LOGIN|ADMIN|JSON))
	r.POST("/region", Preload(Set[Region], LOGIN|ADMIN|JSON))
	r.POST("/event", Preload(Set[Event], LOGIN|ADMIN|JSON))
	r.POST("/history", Preload(Set[History], LOGIN|ADMIN|JSON))
	r.POST("/resource", Preload(Set[Resource], LOGIN|ADMIN|JSON))
	r.POST("/notice", Preload(Set[Notice], LOGIN|ADMIN|JSON))
	r.POST("/sensor", Preload(Set[Sensor], LOGIN|ADMIN|JSON))
	r.POST("/analysis", Preload(Set[Analysis], LOGIN|ADMIN|JSON))

	r.DELETE("/users", Preload(Delete[User], LOGIN|ADMIN|JSON))
	r.DELETE("/regions", Preload(Delete[Region], LOGIN|ADMIN|JSON))
	r.DELETE("/events", Preload(Delete[Event], LOGIN|ADMIN|JSON))
	r.DELETE("/histories", Preload(Delete[History], LOGIN|ADMIN|JSON))
	r.DELETE("/resources", Preload(Delete[Resource], LOGIN|ADMIN|JSON))
	r.DELETE("/notices", Preload(Delete[Notice], LOGIN|ADMIN|JSON))
	r.DELETE("/sensors", Preload(Delete[Sensor], LOGIN|ADMIN|JSON))
	r.DELETE("/analyses", Preload(Delete[Analysis], LOGIN|ADMIN|JSON))

	return r
}
