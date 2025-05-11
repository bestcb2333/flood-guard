package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListHistoryDTO struct {
	RegionID uint `form:"regionId"`
	ListDTO
}

func AddHistoryRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/histories", CreateListHandler[History](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListHistoryDTO{0, ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListHistoryDTO) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			return query
		},
	))

	r.POST("/histories", CreateAddHandler[History](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{JSON: true},
		},
		new(HistoryDTO),
		func(data *History, u *User, dto *HistoryDTO) *History {
			return data
		},
	))

	r.DELETE("/histories", CreateDeleteHandler[History](&PreloaderConfig{
		PreloaderBaseConfig: pbc,
		Bind:                BindConfig{Query: true},
	}))
}

func AddHistoryTrendRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/trends", func(c *gin.Context) {

		var trend []Region
		if err := pbc.DB.Preload("Histories", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "created_at", "region_id", "rainfall", "waterlevel").Order("created_at asc")
		}).Select("id", "name", "updated_at").Find(&trend).Error; err != nil {
			c.JSON(500, Resp("查询失败", err, nil))
			return
		}

		c.JSON(200, Resp("查询成功", nil, trend))
	})
}
