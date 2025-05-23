package main

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
)

type UploadEventDTO struct {
	FH          *multipart.FileHeader `form:"file"`
	StartTime   time.Time             `form:"startTime"`
	EndTime     time.Time             `form:"endTine"`
	Severity    string                `form:"severity"`
	Description string                `form:"description"`
}

func AddUploadEventRoutes(r *gin.Engine, pbc *PreloaderBaseConfig) {

	r.POST("/regions", Preload(
		&PreloaderConfig{},
		&UploadEventDTO{},
		func(c *gin.Context, u *User, r *UploadEventDTO) {

			geo, err := ReadFileHeader[orb.Point](r.FH)
			if err != nil {
				c.JSON(400, Resp("数据格式有误", err, nil))
				return
			}

			if err := pbc.DB.Save(&Event{
				EventDTO: &EventDTO{
					StartTime:   r.StartTime,
					EndTime:     &r.EndTime,
					Severity:    r.Severity,
					Coordinate:  *geo,
					Description: r.Description,
				},
				UserID: &u.ID,
			}).Error; err != nil {
				c.JSON(500, Resp("数据存储失败", err, nil))
				return
			}

			c.JSON(200, Resp("数据上传成功", nil, nil))
		},
	))
}
