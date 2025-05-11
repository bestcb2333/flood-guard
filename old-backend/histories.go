package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetHistoryData(c *gin.Context, _ *User, r struct {
	Region   *uint `form:"region"`
	PageSize int   `form:"pageSize"`
	Page     int   `form:"page"`
}) {

	query := DB.Preload("Region", Select("id", "name")).Model(new(History))
	if r.Region != nil {
		query = query.Where("region_id = ?", r.Region)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(500, Resp("获取总数失败", err, nil))
		return
	}

	var data []History
	if r.PageSize == 0 {
		r.PageSize = 10
	}
	query = query.Limit(r.PageSize)
	if r.Page == 0 {
		r.Page = 1
	}
	query = query.Offset((r.Page - 1) * r.PageSize)
	if err := query.Find(&data).Error; err != nil {
		c.JSON(500, Resp("获取总数失败", err, nil))
		return
	}

	c.JSON(200, Resp("获取成功", nil, gin.H{
		"data":  data,
		"total": total,
	}))
}

func GetHistoryTrend(c *gin.Context) {

	var trend []Region
	if err := DB.Preload("Histories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "created_at", "region_id", "rainfall", "waterlevel").Order("created_at asc")
	}).Select("id", "name", "updated_at").Find(&trend).Error; err != nil {
		c.JSON(500, Resp("查询失败", err, nil))
		return
	}

	c.JSON(200, Resp("查询成功", nil, trend))
}
