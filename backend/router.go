package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB, config *Config) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	pbc := &PreloaderBaseConfig{
		DB:     db,
		JWTKey: config.JWTKey,
	}

	r.GET("/ping", Ping)
	r.GET("/captcha", GetCaptcha)

	AddSignupRoutes(r, pbc)
	AddFactorRoute(r, pbc)
	AddLoginRoutes(r, pbc)
	AddSendEmailRoutes(r, pbc, &config.SMTP)
	AddHistoryTrendRoutes(r, pbc)

	AddUserRoutes(r, pbc)
	AddRegionRoutes(r, pbc)
	AddEventRoutes(r, pbc)
	AddHistoryRoutes(r, pbc)
	AddNoticeRoutes(r, pbc)
	AddSensorRoutes(r, pbc)
	AddResourceRoutes(r, pbc)

	return r
}
