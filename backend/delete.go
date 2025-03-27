package main

import (
	"github.com/gin-gonic/gin"
)

func Delete[T any](c *gin.Context, u *User, r []uint) {

	if err := DB.Model(new(T)).Where(
		"id IN ?", r,
	).Error; err != nil {
		c.JSON(500, Resp("数据删除失败", err, nil))
		return
	}

	c.JSON(200, Resp("操作成功", nil, nil))
}
