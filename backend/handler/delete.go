package handler

import (
	"fmt"
	"strings"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var deleteActionMap = map[string]any{
	"user":         database.User{},
	"region":       database.Region{},
	"floodevent":   database.FloodEvent{},
	"historydata":  database.HistoryData{},
	"notice":       database.Notice{},
	"comment":      database.Comment{},
	"sensor":       database.Sensor{},
	"sensorstatus": database.SensorStatus{},
}

func DeleteRecord(c *gin.Context) {

	user, err := VerifyJwt(c)
	if err != nil {
		util.Error(c, 400, "无法验证你的用户身份", nil)
		return
	}

	path := c.Param("path")
	if viper.GetString(
		fmt.Sprintf("PERMISSION_%s", strings.ToUpper(path)),
	) != "true" {
		if !user.Admin {
			util.Error(c, 400, "只有管理员才能删除数据", nil)
			return
		}
	}

	var ids []uint
	if err := util.ParseJSON(c, &ids); err != nil {
		util.Error(c, 400, "你提供的ID列表不正确", err)
		return
	}

	action := deleteActionMap[path]

	if err := DB.Delete(&action, "id IN ?", ids).Error; err != nil {
		util.Error(c, 500, "数据删除失败", err)
		return
	}

	util.Info(c, 200, "操作成功", nil)
}
