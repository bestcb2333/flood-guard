package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Field struct {
	Key string `json:"key"`
	FieldMeta
}

func Get[T any](c *gin.Context) {

	query := DB.Model(new(T))
	params := c.Request.URL.Query()
	tableName := c.Request.URL.Path[1:]
	tableMeta := Metadata[tableName]

	// 加载相关的表
	if preloads := tableMeta.Preloads; preloads != nil {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	if startTime := params.Get("start_time"); startTime != "" {
		startTime, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			c.JSON(400, Resp("你提供的时间格式不正确", err, nil))
			return
		}
		query = query.Where("created_at >= ?", startTime)
	}

	if endTime := params.Get("end_time"); endTime != "" {
		endTime, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			c.JSON(400, Resp("你提供的时间格式不正确", err, nil))
			return
		}
		query = query.Where("created_at <= ?", endTime)
	}

	if order := params.Get("order"); order != "" {
		query = query.Order(order)
	}

	for key, value := range params {
		if fieldMeta, exists := tableMeta.Fields[key]; exists {
			switch fieldMeta.Type {
			case "string":
				query = query.Scopes(StringScope(key, value[0]))
			case "int":
				query = query.Scopes(IntScope(key, value[0]))
			case "bool":
				query = query.Scopes(BoolScope(key, value[0]))
			}
		}
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		c.JSON(500, Resp("获取记录总数失败", err, nil))
		return
	}

	var data []T
	if err := query.Find(&data).Error; err != nil {
		c.JSON(500, Resp("查询数据失败", err, nil))
		return
	}

	c.JSON(200, Resp("获取成功", nil, gin.H{
		"fields": tableMeta.Fields,
		"count":  count,
		"data":   data,
	}))
}

func StringScope(
	condition, value string,
) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			db = db.Where(fmt.Sprintf("%s = ?", condition), value)
		}
		return db
	}
}

func IntScope(
	condition, value string,
) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if len(value) < 3 {
			return db
		}

		operator := value[:2]
		val, err := strconv.Atoi(value[2:])
		if err != nil {
			return db
		}

		switch operator {
		case "lt":
			return db.Where(condition+" < ?", val)
		case "gt":
			return db.Where(condition+" > ?", val)
		default:
			return db
		}
	}
}

func BoolScope(
	condition, value string,
) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition+" = ?", value)
	}
}
