package handler

import (
	"io"

	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
)

func PostMidWare(c *gin.Context) {

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		util.Error(c, 400, "无法读取你的请求", nil)
		return
	}

	c.Set("requestBody", bodyBytes)
}
