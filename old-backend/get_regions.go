package main

import "github.com/gin-gonic/gin"

func GetRegions(c *gin.Context) {

	var data []Region
	if err := DB.Find(&data).Error; err != nil {
		c.JSON(500, Resp("查询失败", err, nil))
		return
	}

	c.JSON(200, Resp("查询成功", nil, data))
}
