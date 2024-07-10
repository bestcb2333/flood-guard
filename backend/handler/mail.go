package handler

import (
	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func GetMail(c *gin.Context) {

	email := c.Query("email")

	if err := DB.First(
		new(db.User), "email = ?", email,
	).Error; err == nil {
		util.Error(c, 400, "这个邮箱已经被注册过了", nil)
		return
	}
}
