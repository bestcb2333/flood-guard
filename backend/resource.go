package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListResourceDTO struct {
	Type   string `form:"type"`
	Region uint   `form:"region"`
	ListDTO
}

func AddResourceRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/resources", CreateListHandler[Resource](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListResourceDTO{"", 0, ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListResourceDTO) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			if r.Type != "" {
				query = query.Where("type = ?", r.Type)
			}
			if r.Region != 0 {
				query = query.Where("region_id = ?", r.Region)
			}
			return query
		},
	))

	r.POST("/resources", CreateAddHandler(
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{JSON: true},
		},
		new(ResourceDTO),
		func(data *Resource, u *User, dto *ResourceDTO) *Resource {
			data.ResourceDTO = *dto
			return data
		},
	))

	r.DELETE("/resources", CreateDeleteHandler[Resource](&PreloaderConfig{
		PreloaderBaseConfig: pbc,
		Bind:                BindConfig{Query: true},
	}))
}
