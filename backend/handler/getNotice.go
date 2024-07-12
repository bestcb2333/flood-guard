package handler

import (
	"strconv"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func GetNotice(c *gin.Context) {

	query := DB.Model(new(database.Notice))

	if title := c.Query("title"); title != "" {
		query = query.Where("title = ?", title)
	}

	if author, err := strconv.Atoi(c.Query("author")); err != nil {
		query = query.Where("author = ?", author)
	}

	if limit, err := strconv.Atoi(c.Query("limit")); err != nil {
		query = query.Limit(limit)
	}

	if offset, err := strconv.Atoi(c.Query("offset")); err != nil {
		query = query.Offset(offset)
	}

	var data []database.Notice
	if err := query.Find(&data).Error; err != nil {
		util.HandleQueryErr(c, "为找到符合条件的记录", err)
		return
	}

	util.Info(c, 200, "查询成功", data)
}
