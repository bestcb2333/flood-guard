package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

var SelectActionMap = map[string]map[string]string{
	"sensor": {
		"string": "status name",
		"int":    "region",
	},
	"sensorstatus": {
		"from":   "time",
		"to":     "time",
		"int":    "sensor",
		"string": "status",
	},
	"floodevent": {
		"from":   "start_time",
		"to":     "end_time",
		"int":    "region uploader",
		"string": "severity",
	},
	"historydata": {
		"from": "record_time",
		"to":   "record_time",
		"int":  "region",
	},
	"notice": {
		"string": "title",
		"int":    "author",
	},
	"user": {
		"bool": "admin",
	},
}

func SelectRecord(c *gin.Context) {

	action := c.Request.URL.Path[5:]
	actionMap := SelectActionMap[action]
	query := DB
	var data any

	switch action {
	case "sensor":
		data = make([]database.Sensor, 0)
		query = query.Model(new(database.Sensor))
	case "sensorstatus":
		data = make([]database.SensorStatus, 0)
		query = query.Model(new(database.SensorStatus))
	case "floodevent":
		data = make([]database.FloodEvent, 0)
		query = query.Model(new(database.FloodEvent))
	case "historydata":
		data = make([]database.HistoryData, 0)
		query = query.Model(new(database.HistoryData))
	case "notice":
		data = make([]database.Notice, 0)
		query = query.Model(new(database.Notice))
	case "region":
		data = make([]database.Region, 0)
		query = query.Model(new(database.Region))
	case "user":
		data = make([]database.User, 0)
		query = query.Model(new(database.User))
	default:
		util.Error(c, 400, "无效的查询类型", nil)
		return
	}

	if actionMap != nil {

		if value, ok := actionMap["from"]; ok {
			if from, err := time.Parse(
				"2006-01-02 15:04:05", c.Query(value),
			); err == nil {
				query = query.Where(value+" >= ?", from)
			}
		}

		if value, ok := actionMap["to"]; ok {
			if to, err := time.Parse(
				"2006-01-02 15:04:05", c.Query(value),
			); err == nil {
				query = query.Where(value+" <= ?", to)
			}
		}

		if intSlice, ok := actionMap["int"]; ok {
			for _, value := range strings.Split(intSlice, " ") {
				if data, err := strconv.Atoi(c.Query(value)); err != nil {
					query = query.Where(value+" = ?", data)
				}
			}
		}

		if stringSlice, ok := actionMap["string"]; ok {
			for _, value := range strings.Split(stringSlice, " ") {
				if data := c.Query(value); data != "" {
					query = query.Where(value+" = ?", data)
				}
			}
		}

		if boolSlice, ok := actionMap["bool"]; ok {
			for _, value := range strings.Split(boolSlice, " ") {
				if data := c.Query(value); data == "true" {
					query = query.Where(value+" = ?", 1)
				}
			}
		}
	}

	if rawLimit := c.Query("limit"); rawLimit != "" {
		if limit, err := strconv.Atoi(rawLimit); err != nil {
			util.Error(c, 400, "你提供的限制数量不正确", err)
			return
		} else {
			query = query.Limit(limit)
		}
	}

	if rawOffset := c.Query("offset"); rawOffset != "" {
		if offset, err := strconv.Atoi(rawOffset); err != nil {
			util.Error(c, 400, "你提供的起始位置不正确", err)
			return
		} else {
			query = query.Offset(offset)
		}
	}

	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "无法找到符合条件的记录", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
