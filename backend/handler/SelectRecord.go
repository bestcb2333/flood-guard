package handler

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

var SelectActionMap = map[string]map[string]string{
	"user": {
		"bool": "admin",
	},
	"floodevent": {
		"from":   "start_time",
		"to":     "end_time",
		"int":    "region uploader",
		"string": "severity",
	},
	"historydata": {
		"from": "start_time",
		"to":   "end_time",
		"int":  "region",
	},
	"notice": {
		"string": "title",
		"int":    "author",
	},
	"comment": {
		"string": "related",
		"int":    "author",
	},
	"sensor": {
		"string": "status name",
		"int":    "region",
	},
	"sensorstatus": {
		"from":   "start_time",
		"to":     "end_time",
		"int":    "sensor",
		"string": "status",
	},
}

var SelectTableMap = map[string]any{
	"user":         make([]database.User, 1),
	"region":       make([]database.Region, 1),
	"floodevent":   make([]database.FloodEvent, 1),
	"historydata":  make([]database.HistoryData, 1),
	"notice":       make([]database.Notice, 1),
	"comment":      make([]database.Comment, 1),
	"sensor":       make([]database.Sensor, 1),
	"sensorstatus": make([]database.SensorStatus, 1),
}

type Field struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

func SelectRecord(c *gin.Context) {

	action := c.Param("path")
	actionMap := SelectActionMap[action]
	v := reflect.ValueOf(SelectTableMap[action])
	if v.Kind() != reflect.Slice {
		util.Error(c, 400, "路径不存在", nil)
		return
	}
	query := DB.Model(v.Index(0).Interface())
	structType := v.Type().Elem()
	dataReflect := reflect.MakeSlice(
		reflect.SliceOf(structType), 0, 0,
	)
	reflect.Copy(dataReflect, v)
	data := dataReflect.Interface()

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
				if data, err := strconv.Atoi(c.Query(value)); err == nil {
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

	if order := c.Query("order"); order != "" {
		query = query.Order(order)
	}

	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "无法找到符合条件的记录", err)
		return
	}

	fieldSlice := make([]Field, 0)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		title := field.Tag.Get("gorm")
		if strings.HasPrefix(title, "comment") {
			fieldSlice = append(fieldSlice, Field{
				Key:   strings.ToLower(field.Name),
				Title: title[8:],
			})
		}
	}

	c.AbortWithStatusJSON(200, gin.H{
		"msg":   "查询成功",
		"count": dataReflect.Len(),
		"field": fieldSlice,
		"data":  data,
	})
}
