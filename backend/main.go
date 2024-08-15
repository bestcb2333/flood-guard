package main

import (
	"github.com/bestcb2333/FloodGuard/config"
	"github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/handler"
)

func main() {
	config.Init()
	database.Init()
	handler.Init()
	Router()
}
