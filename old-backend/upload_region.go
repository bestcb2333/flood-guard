package main

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
)

type FeatureCollection[T any] struct {
	Type     string       `json:"type"`
	Features []Feature[T] `json:"features"`
}

type Feature[T any] struct {
	Type       string         `json:"type"`
	Geometry   Geometry[T]    `json:"geometry"`
	Properties map[string]any `json:"properties"`
}

type Geometry[T any] struct {
	Type        string `json:"type"`
	Coordinates T      `json:"coordinates"`
}

func UploadRegion(c *gin.Context, u *User, r struct {
	FH          *multipart.FileHeader `form:"file"`
	Name        string                `form:"name"`
	Description string                `form:"description"`
}) {

	if err := DB.Unscoped().Where("1=1").Delete(new(Region)).Error; err != nil {
		c.JSON(500, Resp("regions表清空失败", err, nil))
		return
	}

	geo, err := ReadFileHeader[FeatureCollection[orb.MultiPolygon]](r.FH)
	if err != nil {
		c.JSON(400, Resp("文件格式有误", err, nil))
		return
	}

	for _, feature := range geo.Features {
		region := &Region{
			Coordinate: feature.Geometry.Coordinates,
			Name:       feature.Properties["name"].(string),
		}
		if err := DB.Save(region).Error; err != nil {
			c.JSON(500, Resp("数据存储失败", err, nil))
			return
		}
	}

	c.JSON(200, Resp("数据存储完成", nil, nil))
}
