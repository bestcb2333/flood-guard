package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type Paged interface {
	GetListDTO() *ListDTO
}

func ListHandler[T any, DTO Paged](
	def *DTO,
	queryFunc func(db *gorm.DB, u *User, dto *DTO) *gorm.DB,
) gin.HandlerFunc {
	return Preload(
		&PreloaderConfig{DB, Config.JWTKey, binding.Query, nil, Public, nil},
		def,
		func(c *gin.Context, u *User, r *DTO) {
			query := DB.Model(new(T))
			query = queryFunc(query, u, r)

			var total int64
			if err := query.Count(&total).Error; err != nil {
				c.JSON(500, Resp("获取总数失败", err, nil))
				return
			}

			dto := (*r).GetListDTO()
			query = query.Limit(dto.PageSize).Offset((dto.Page - 1) * dto.PageSize)

			var data []T
			if err := query.Find(&data).Error; err != nil {
				c.JSON(500, Resp("数据查询失败", err, nil))
				return
			}

			c.JSON(200, Resp("数据查询成功", nil, gin.H{
				"data":  data,
				"total": total,
			}))
		},
	)
}
