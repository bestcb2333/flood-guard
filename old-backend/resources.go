package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GetResourcesReq struct {
	Type   string `form:"type"`
	Region uint   `form:"region"`
	ListDTO
}

var GetResources = Preload(
	&PreloaderConfig{DB, Config.JWTKey, binding.Query, nil, Public, nil},
	&GetResourcesReq{"", 0, ListDTO{1, 10}},
	func(c *gin.Context, _ *User, r *GetResourcesReq) {

		query := DB.Preload("Region", Select("id", "name")).Model(new(Resource))
		if r.Type != "" {
			query = query.Where("type = ?", r.Type)
		}
		if r.Region != 0 {
			query = query.Where("region_id = ?", r.Region)
		}

		var total int64
		if err := query.Count(&total).Error; err != nil {
			c.JSON(500, Resp("获取总数失败", err, nil))
			return
		}

		var data []Resource
		query = query.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize)
		if err := query.Find(&data).Error; err != nil {
			c.JSON(500, Resp("数据查询失败", err, nil))
			return
		}

		c.JSON(200, Resp("数据查询成功", nil, gin.H{
			"data":  data,
			"total": total,
		}))

	},
)

var AddResource = Preload(
	&PreloaderConfig{DB, Config.JWTKey, nil, binding.JSON, Login, nil},
	&ResourceDTO{},
	func(c *gin.Context, u *User, r *ResourceDTO) {

		resource := Resource{ResourceDTO: *r}

		if err := DB.Create(&resource).Error; err != nil {
			c.JSON(500, Resp("资源创建失败", err, nil))
			return
		}

		c.JSON(200, Resp("资源创建成功", nil, &resource))
	},
)
