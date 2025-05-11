package main

import "github.com/gin-gonic/gin"

func GetNotices(c *gin.Context, _ *User, r struct {
	Page     uint `form:"page"`
	PageSize int  `form:"page"`
}) {

	query := DB.Preload("User").Model(new(Notice))

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(500, Resp("获取数量失败", err, nil))
		return
	}

	var data []Notice
	if r.PageSize != 0 {
		r.PageSize = 10
	}
	if err := query.Order("created_at desc").Find(&data).Error; err != nil {
		c.JSON(500, Resp("查询失败", err, nil))
		return
	}

	c.JSON(200, Resp("查询成功", nil, gin.H{
		"count": total,
		"data":  data,
	}))
}
