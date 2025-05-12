package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListSensorReq struct {
	RegionID  uint `form:"region_id"`
	Available bool `form:"available"`
	p.PageConfig
}

func AddSensorRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/sensors", p.CreateListHandler[Sensor](
		&p.Config[ListSensorReq]{
			Base: bc,
			DefReq: ListSensorReq{
				RegionID:  0,
				Available: false,
				PageConfig: p.PageConfig{
					Page:     1,
					PageSize: 10,
				},
			},
		},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListSensorReq) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			if r.Available {
				query = query.Where("available = true")
			}
			return query
		},
	))

	r.POST("/sensors", p.CreateAddHandler[Sensor](
		&p.Config[SensorDTO]{
			Base: bc,
		},
		func(c *gin.Context, u *User, dto *SensorDTO) *Sensor {
			data := new(Sensor)
			return data
		},
	))

	r.DELETE("/sensors", p.CreateDeleteHandler[Sensor, User](
		&p.Config[p.DelReq]{
			Base: bc,
		},
	))
}
