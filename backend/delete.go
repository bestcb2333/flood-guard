package main

import (
	"github.com/gin-gonic/gin"
)

func CreateDeleteHandler[T any](
	pc *PreloaderConfig,
) gin.HandlerFunc {
	pc.Bind = BindConfig{Query: true}
	return Preload(
		pc,
		&DeleteDTO{},
		func(c *gin.Context, u *User, r *DeleteDTO) {

			if err := pc.DB.Where("id IN ?", r.IDs).Delete(new(T)).Error; err != nil {
				c.JSON(500, Resp("删除失败", err, nil))
				return
			}

			c.JSON(200, Resp("删除成功", nil, nil))
		},
	)
}
