package handler

import (
	"strconv"
	"time"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func GetFloodEvent(c *gin.Context) {

	query := DB.Model(new(database.FloodEvent))

	if region, err := strconv.Atoi(c.Query("region")); err != nil {
		query = query.Where("region_id = ?", region)
	}

	if from, err := time.Parse(
		"2006-01-02 15:04:05", c.Query("from"),
	); err != nil {
		query = query.Where("start_time >= ?", from)
	}

	if to, err := time.Parse(
		"2006-01-02 15:04:05", c.Query("end"),
	); err != nil {
		query = query.Where("end_time <= ?", to)
	}

	if uploader, err := strconv.Atoi(c.Query("uploader")); err != nil {
		query = query.Where("uploader = ?", uploader)
	}

	if severity := c.Query("severity"); severity != "" {
		query = query.Where("severity = ?", severity)
	}

	var data []database.FloodEvent
	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "查不到符合条件的记录", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
