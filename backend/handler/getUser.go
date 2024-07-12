package handler

import (
	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	query := DB.Model(new(database.User))

	if admin := c.Query("admin"); admin == "true" {
		query = query.Where("admin = ?", 1)
	}

	var data []database.User
	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "无法找到符合条件的记录", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
