package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListEventDTO struct {
	RegionID uint `json:"regionId"`
	ListDTO
}

func AddEventRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/events", CreateListHandler[Event](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListEventDTO{0, ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListEventDTO) *gorm.DB {
			query = query.Preload("User", Select("id", "name")).Preload("Region", Select("id", "name"))
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			return query
		},
	))

	r.POST("/events", CreateAddHandler[Event](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{JSON: true},
		},
		&EventDTO{},
		func(data *Event, u *User, dto *EventDTO) *Event {
			data.UserID = &u.ID
			return data
		},
	))

	r.DELETE("/events", CreateDeleteHandler[Event](&PreloaderConfig{
		PreloaderBaseConfig: pbc,
		Bind:                BindConfig{Query: true},
	}))
}
