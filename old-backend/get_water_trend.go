package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWaterTrend(c *gin.Context) {

	var data []Region
	if err := DB.Preload("Histories", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC").Limit(7)
	}).Find(&data).Error; err != nil {
		c.JSON(500, Resp("数据获取失败", err, nil))
		return
	}

	c.JSON(200, Resp("查询成功", nil, data))
}
