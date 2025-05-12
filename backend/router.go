package main

import (
	"time"

	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
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

	bc := &p.BaseConfig{
		DB:            db,
		JWTKey:        config.JWTKey,
		JWTExpHours:   24 * time.Hour,
		UserTableName: "users",
		AdminColName:  "admin",
		Resp: func(message string, err error, data any) gin.H {
			var errStr *string
			if err != nil {
				str := err.Error()
				errStr = &str
			}
			return gin.H{
				"message": message,
				"error":   errStr,
				"data":    data,
			}
		},
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
	AddEventRoutes(r, bc)
	AddHistoryRoutes(r, bc)
	AddNoticeRoutes(r, bc)
	AddSensorRoutes(r, bc)
	AddResourceRoutes(r, bc)

	return r
}
