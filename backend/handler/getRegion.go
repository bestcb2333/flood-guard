package handler

import (
	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func GetRegion(c *gin.Context) {

	var data []database.Region
	if err := DB.Find(&data).Error; err != nil {
		util.Error(c, 500, "查询失败", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
