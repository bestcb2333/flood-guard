package handler

import (
	"time"

	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

// 查询历史记录
func GetHistoryData(c *gin.Context) {

	query := DB.Model(new(db.HistoryData)).Joins("Region").Preload("Region")

	if region := c.Query("region"); region != "" {
		query = query.Where("region.name = ?", region)
	}

	if startTime, err := time.Parse(
		"2006-01-02 15:04:05", c.Query("start"),
	); err != nil {
		query = query.Where("record_time >= ?", startTime)
	}

	if endTime, err := time.Parse(
		"2006-01-02 15:04:05", c.Query("end"),
	); err != nil {
		query = query.Where("record_time <= ?", endTime)
	}

	var data []db.HistoryData
	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "找不到符合条件的记录", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
