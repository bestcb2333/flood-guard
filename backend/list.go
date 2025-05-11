package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Paged interface {
	GetListDTO() *ListDTO
}

func CreateListHandler[T any, DTO Paged](
	pc *PreloaderConfig,
	def *DTO,
	queryFunc func(q *gorm.DB, c *gin.Context, u *User, r *DTO) *gorm.DB,
) gin.HandlerFunc {
	return Preload(
		pc,
		def,
		func(c *gin.Context, u *User, r *DTO) {

			query := pc.DB.Model(new(T))
			query = queryFunc(query, c, u, r)
			if query == nil {
				return
			}

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
