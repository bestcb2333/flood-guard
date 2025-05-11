package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb/geojson"
)

func GetMap(c *gin.Context, _ *User, r struct {
	EventID *uint `form:"event_id"`
}) {

	query := DB.Model(new(Region))
	var regions []Region
	fc := geojson.NewFeatureCollection()
	if err := query.Find(&regions).Error; err != nil {
		c.JSON(500, Resp("区域查询失败", err, nil))
		return
	}

	for _, region := range regions {
		regionFeature := geojson.NewFeature(region.Coordinate)
		regionFeature.Properties["name"] = region.Name
		regionFeature.Properties["description"] = region.Description
		fc.Append(regionFeature)
	}

	if r.EventID != nil {

		var event Event
		if err := DB.Last(&event).Error; err != nil {
			c.JSON(500, Resp("事件信息获取失败", err, nil))
			return
		}
		feature := geojson.NewFeature(event.Coordinate)
		feature.Properties["startTime"] = event.StartTime
		feature.Properties["endTime"] = event.EndTime
		feature.Properties["severity"] = event.Severity
		fc.Append(feature)
	}

	c.JSON(200, Resp("查询成功", nil, fc))
}
