package main

import (
	"log"
)

func main() {

	var config Config
	if err := LoadConfigFromEnv(&config); err != nil {
		log.Fatalf("配置读取失败：%w\n", err)
	}

	db, err := InitDB(&config)
	if err != nil {
		log.Fatalf("项目初始化失败: %v\n", err.Error())
	}

	r := GetRouter(db, &config)

	if config.SSL.Enable {
		r.RunTLS(
			":"+config.Port,
			config.SSL.Certificate,
			config.SSL.Key,
		)
	} else {
		r.Run(":" + config.Port)
	}
}
