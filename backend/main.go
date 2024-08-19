package main

import (
	"log"

	"github.com/bestcb2333/FloodGuard/config"
	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/handler"
	"github.com/bestcb2333/FloodGuard/util"
)

func main() {
	config.Init()
	if err := database.Init(); err != nil {
		log.Fatalf("程序初始化失败: %v\n", err.Error())
	}
	if err := util.Init(); err != nil {
		log.Fatal("程序初始化失败: %v\n", err.Error())
	}
	handler.Init()
	Router()
}
