package main

import (
	"github.com/gin-gonic/gin"
)

func CreateAddHandler[T any, DTO any](
	pc *PreloaderConfig,
	def *DTO,
	handlerFunc func(data *T, u *User, dto *DTO) *T,
) gin.HandlerFunc {
	return Preload(
		pc,
		def,
		func(c *gin.Context, u *User, r *DTO) {

			data := new(T)

			data = handlerFunc(data, u, r)

			if err := pc.DB.Create(data).Error; err != nil {
				c.JSON(500, Resp("资源创建失败", err, nil))
				return
			}

			c.JSON(200, Resp("数据创建成功", nil, data))

		},
	)
}
