package main

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context, _ *User, r struct {
	Region *uint `form:"region"`
	Limit  *int  `form:"limit"`
}) {

	query := DB.Preload("Region", Select("id", "name")).Model(new(User)).Select("id", "name", "region_id")
	if r.Region != nil {
		query = query.Where("region_id = ?", r.Region)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		c.JSON(500, Resp("获取总数失败", err, nil))
		return
	}

	var data []User
	if r.Limit != nil {
		query = query.Limit(*r.Limit)
	}
	if err := query.Find(&data).Error; err != nil {
		c.JSON(500, Resp("数据查询失败", err, nil))
		return
	}

	c.JSON(200, Resp("数据查询成功", nil, gin.H{
		"data":  data,
		"count": count,
	}))
}
