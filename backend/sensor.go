package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListSensorDTO struct {
	RegionID uint `form:"regionId"`
	ListDTO
}

func AddSensorRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/sensors", CreateListHandler[Sensor](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},

		&ListSensorDTO{0, ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, dto *ListSensorDTO) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if dto.RegionID != 0 {
				query = query.Where("region_id = ?", dto.RegionID)
			}

			return query
		},
	))

	r.POST("/sensors", CreateAddHandler[Sensor](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
		},
		new(SensorDTO),
		func(data *Sensor, u *User, dto *SensorDTO) *Sensor {
			return data
		},
	))

	r.DELETE("/sensors", CreateDeleteHandler[Sensor](&PreloaderConfig{
		PreloaderBaseConfig: pbc,
		Bind:                BindConfig{Query: true},
	}))
}
