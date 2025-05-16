package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListEventReq struct {
	RegionID uint   `form:"region_id"`
	Severity string `form:"severity"`
	Current  bool   `form:"current"`
	p.PageConfig
}

func AddEventRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/events", p.CreateListHandler[Event](
		&p.Config[ListEventReq]{
			Base: bc,
			DefReq: ListEventReq{
				RegionID: 0,
				Severity: "",
				Current:  false,
				PageConfig: p.PageConfig{
					Page:     1,
					PageSize: 10,
				},
			},
		},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListEventReq) *gorm.DB {
			query = query.Preload("User", Select("id", "name")).Preload("Region", Select("id", "name"))
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			if r.Severity != "" {
				query = query.Where("severity = ?", r.Severity)
			}
			if r.Current {
				query = query.Where("end_time IS NULL")
			}
			return query
		},
	))

	r.POST("/events", p.CreateAddHandler[Event](
		&p.Config[EventDTO]{
			Base:       bc,
			Permission: p.Login,
		},
		func(c *gin.Context, u *User, dto *EventDTO) *Event {
			data := new(Event)
			data.UserID = &u.ID
			data.EventDTO = dto
			return data
		},
	))

	r.DELETE("/events", p.CreateDeleteHandler[Event, User](
		&p.Config[p.DelReq]{
			Base: bc,
		},
	))
}
