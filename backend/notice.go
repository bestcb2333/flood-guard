package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListNoticeDTO struct {
	ListDTO
}

func AddNoticeRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/notices", CreateListHandler[Notice](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListNoticeDTO{ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, dto *ListNoticeDTO) *gorm.DB {
			return query
		},
	))

	r.POST("/notices", CreateAddHandler[Notice](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		new(NoticeDTO),
		func(data *Notice, u *User, dto *NoticeDTO) *Notice {
			data.UserID = &u.ID
			return data
		},
	))

	r.DELETE("/notices", CreateDeleteHandler[Notice](&PreloaderConfig{
		PreloaderBaseConfig: pbc,
		Bind:                BindConfig{Query: true},
	}))

}
