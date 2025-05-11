package main

import "github.com/gin-gonic/gin"

func AuthToken(
	logicFunc func(*User, *gin.Context), preloads ...string,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		user, err := GetUserByJwt(c.GetHeader("Authorization"), preloads...)
		if err != nil {
			c.JSON(400, Resp("读取用户信息失败", err, nil))
			return
		}

		logicFunc(user, c)
	}
}
