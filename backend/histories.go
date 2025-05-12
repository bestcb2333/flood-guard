package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListHistoryReq struct {
	RegionID uint `form:"region_id"`
	p.PageConfig
}

func AddHistoryRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/histories", p.CreateListHandler[History](
		&p.Config[ListHistoryReq]{
			Base: bc,
		},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListHistoryReq) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			return query
		},
	))

	r.POST("/histories", p.CreateAddHandler[History](
		&p.Config[HistoryDTO]{
			Base: bc,
		},
		func(c *gin.Context, u *User, r *HistoryDTO) *History {
			data := new(History)
			return data
		},
	))

	r.DELETE("/histories", p.CreateDeleteHandler[History, User](
		&p.Config[p.DelReq]{
			Base: bc,
		},
	))
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
