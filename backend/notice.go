package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListNoticeReq struct {
	p.PageConfig
}

func AddNoticeRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/notices", p.CreateListHandler[Notice](
		&p.Config[ListNoticeReq]{
			Base: bc,
			DefReq: ListNoticeReq{
				PageConfig: p.PageConfig{
					Page:     1,
					PageSize: 10,
				},
			},
		},
		func(query *gorm.DB, c *gin.Context, u *User, dto *ListNoticeReq) *gorm.DB {
			return query
		},
	))

	r.POST("/notices", p.CreateAddHandler[Notice](
		&p.Config[NoticeDTO]{
			Base:       bc,
			Permission: p.Login,
		},
		func(c *gin.Context, u *User, dto *NoticeDTO) *Notice {
			data := new(Notice)
			data.UserID = &u.ID
			data.NoticeDTO = dto
			return data
		},
	))

	r.DELETE("/notices", p.CreateDeleteHandler[Notice, User](
		&p.Config[p.DelReq]{
			Base: bc,
		},
	))

}
