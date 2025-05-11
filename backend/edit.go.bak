package main

import (
	"github.com/gin-gonic/gin"
)

func Set[T any](c *gin.Context, _ *User, r T) {

	if err := DB.Save(&r).Error; err != nil {
		c.JSON(500, Resp("数据编辑失败", err, nil))
		return
	}

	c.JSON(200, Resp("数据编辑成功", nil, nil))
}
