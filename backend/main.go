package main

import (
	"log"
)

func main() {
	if err := Init(); err != nil {
		log.Fatalf("项目初始化失败: %v\n", err.Error())
	}
	r := Router()
	if Config.SSL.Enable {
		r.RunTLS(
			":"+Config.Port,
			Config.SSL.Certificate,
			Config.SSL.Key,
		)
	} else {
		r.Run(":" + Config.Port)
	}
}
