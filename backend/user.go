package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListUserDTO struct {
	ListDTO
}

func AddUserRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.GET("/users", CreateListHandler[User](
		&PreloaderConfig{
			PreloaderBaseConfig: pbc,
			Bind:                BindConfig{Query: true},
		},
		&ListUserDTO{ListDTO{1, 10}},
		func(query *gorm.DB, c *gin.Context, u *User, r *ListUserDTO) *gorm.DB {
			query = query.Preload("Region", Select("id", "name"))
			return query
		},
	))
}
