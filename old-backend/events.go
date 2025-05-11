package main

import "github.com/gin-gonic/gin"

func GetEvents(c *gin.Context, _ *User, r struct {
	Region   *uint `form:"region"`
	PageSize int   `form:"pageSize"`
	Page     int   `form:"page"`
}) {

	query := DB.Preload("Region", Select("id", "name")).Preload("User", Select("id", "name")).Model(new(Event))
	if r.Region != nil {
		query = query.Where("region_id = ?", r.Region)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(500, Resp("获取总数失败", err, nil))
		return
	}

	var data []Event
	if r.PageSize == 0 {
		r.PageSize = 10
	}
	query = query.Limit(r.PageSize)
	if r.Page == 0 {
		r.Page = 1
	}
	query = query.Offset((r.Page - 1) * r.PageSize)
	if err := query.Find(&data).Error; err != nil {
		c.JSON(500, Resp("数据查询失败", err, nil))
		return
	}

	c.JSON(200, Resp("数据查询成功", nil, gin.H{
		"data":  data,
		"total": total,
	}))
}
