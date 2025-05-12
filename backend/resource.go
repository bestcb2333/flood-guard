package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListResourceReq struct {
	Type      string `form:"type"`
	RegionID  uint   `form:"region_id"`
	Available bool   `form:"available"`
	p.PageConfig
}

func AddResourceRoutes(r *gin.Engine, bc *p.BaseConfig) {

	r.GET("/resources", p.CreateListHandler[Resource](
		&p.Config[ListResourceReq]{
			Base: bc,
			DefReq: ListResourceReq{
				Type:      "all",
				Available: false,
				PageConfig: p.PageConfig{
					Page:     1,
					PageSize: 10,
				},
			},
		},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListResourceReq) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if r.Type != "all" {
				query = query.Where("type = ?", r.Type)
			}
			if r.RegionID != 0 {
				query = query.Where("region_id = ?", r.RegionID)
			}
			if r.Available {
				query = query.Where("available = true")
			}
			return query
		},
	))

	r.POST("/resources", p.CreateAddHandler(
		&p.Config[ResourceDTO]{
			Base: bc,
		},
		func(c *gin.Context, u *User, dto *ResourceDTO) *Resource {
			data := new(Resource)
			data.ResourceDTO = *dto
			return data
		},
	))

	r.DELETE("/resources", p.CreateDeleteHandler[Resource, User](
		&p.Config[p.DelReq]{
			Base: bc,
		},
	))
}
