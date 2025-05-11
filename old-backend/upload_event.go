package main

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
)

func UploadEvent(c *gin.Context, u *User, r struct {
	FH          *multipart.FileHeader `form:"file"`
	StartTime   time.Time             `form:"startTime"`
	EndTime     time.Time             `form:"endTine"`
	Severity    uint                  `form:"severity"`
	Description string                `form:"description"`
}) {

	geo, err := ReadFileHeader[orb.Point](r.FH)
	if err != nil {
		c.JSON(400, Resp("数据格式有误", err, nil))
		return
	}

	if err := DB.Save(&Event{
		StartTime:   r.StartTime,
		EndTime:     r.EndTime,
		UserID:      &u.ID,
		Severity:    r.Severity,
		Coordinate:  *geo,
		Description: r.Description,
	}).Error; err != nil {
		c.JSON(500, Resp("数据存储失败", err, nil))
		return
	}

	c.JSON(200, Resp("数据上传成功", nil, nil))
}
