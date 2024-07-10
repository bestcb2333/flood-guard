package handler

import (
	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func addRegion(c *gin.Context) {

	_, err := VerifyJwt(c)
	if err != nil {
		util.Error(c, 400, "只有管理员才能添加新地区", nil)
		return
	}

	regionName := c.Query("region_name")
	if regionName != "" {
		util.Error(c, 400, "区域名称不能为空", nil)
		return
	}

	if err := DB.Create(&db.Region{Name: regionName}).Error; err != nil {
		util.Error(c, 500, "新地区创建失败", nil)
		return
	}

	util.Info(c, 200, "新地区创建成功", nil)
}
