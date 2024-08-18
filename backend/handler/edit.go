package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func EditRecord(c *gin.Context) {

	user, err := VerifyJwt(c)
	if err != nil {
		util.Error(c, 400, "用户身份验证失败", err)
		return
	}

	path := c.Param("path")

	if viper.GetString(
		fmt.Sprintf("PERMISSION_%s", strings.ToUpper(path)),
	) == "false" {
		if !user.Admin {
			util.Error(c, 400, "你不是管理员，无法执行此操作", nil)
			return
		}
	}

	jsonBytesAny, ok := c.Get("requestBody")
	if !ok {
		util.Error(c, 500, "服务器错误", nil)
		return
	}
	if err := EditRecordFunc(jsonBytesAny.([]byte), path); err != nil {
		util.Error(c, 500, "服务器保存记录失败", err)
		return
	}

	util.Info(c, 200, "操作成功", nil)
}

func EditRecordFunc(data []byte, path string) error {
	switch path {
	case "user":
		object, err := UnmarshalToType[database.User](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "region":
		object, err := UnmarshalToType[database.Region](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "floodevent":
		object, err := UnmarshalToType[database.FloodEvent](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "historydata":
		object, err := UnmarshalToType[database.HistoryData](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "notice":
		object, err := UnmarshalToType[database.Notice](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "comment":
		object, err := UnmarshalToType[database.Comment](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "sensor":
		object, err := UnmarshalToType[database.Sensor](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	case "sensorstatus":
		object, err := UnmarshalToType[database.SensorStatus](data)
		if err != nil {
			return err
		}
		return DB.Save(object).Error
	default:
		return errors.New("未知的类型")
	}
}

func UnmarshalToType[T any](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return new(T), err
	}
	return &result, nil
}
