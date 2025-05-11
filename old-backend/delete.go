package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func OldDelete[T any](c *gin.Context, u *User, r []uint) {

	if err := DB.Model(new(T)).Where(
		"id IN ?", r,
	).Error; err != nil {
		c.JSON(500, Resp("数据删除失败", err, nil))
		return
	}

	c.JSON(200, Resp("操作成功", nil, nil))
}

func DeleteHandler[T any]() gin.HandlerFunc {
	return Preload(
		&PreloaderConfig{DB, Config.JWTKey, binding.Query, nil, Login, nil},
		&DeleteDTO{},
		func(c *gin.Context, u *User, r *DeleteDTO) {

			if err := DB.Where("id IN ?", r.IDs).Delete(new(T)).Error; err != nil {
				c.JSON(500, Resp("删除失败", err, nil))
				return
			}

			c.JSON(200, Resp("删除成功", nil, nil))
		},
	)
}
