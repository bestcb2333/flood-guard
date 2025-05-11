package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func AuthEmail(c *gin.Context) {

	email := c.GetHeader("X-Email-Number")
	code := c.GetHeader("X-Email-Code")

	EmailCodeData.Mu.Lock()
	value, exist := EmailCodeData.Data[email]
	EmailCodeData.Mu.Unlock()

	if !exist {
		c.AbortWithStatusJSON(400, Resp("你还没有申请验证码", nil, nil))
		return
	}

	if time.Now().After(value.Expiration) {
		c.AbortWithStatusJSON(400, Resp("验证码已过期", nil, nil))
		return
	}

	if value.Code != code {
		c.AbortWithStatusJSON(400, Resp("验证码不正确", nil, nil))
		return
	}
}
