package main

import "github.com/gin-gonic/gin"

func Router() {

	r := gin.Default()
	r.RunTLS()
}
