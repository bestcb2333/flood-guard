package main

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {

	c.JSON(200, Resp("Success", nil, nil))
}
