package handler

import (
	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

var DeleteActionMap = map[string]any{
	"delete-user":        database.User{},
	"delete-region":      database.Region{},
	"delete-floodevent":  database.FloodEvent{},
	"delete-historydata": database.HistoryData{},
	"delete-notice":      database.Notice{},
}

func DeleteRecord(c *gin.Context) {

	user, err := VerifyJwt(c)
	if err != nil {
		util.Error(c, 400, "无法验证你的用户身份", nil)
		return
	}

}
