package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListRegionDTO struct {
	ListDTO
}

func AddRegionRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/regions", CreateListHandler[Region](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListRegionDTO{ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListRegionDTO) *gorm.DB {
			return query
		},
	))
}
