package handler

import (
	"github.com/bestcb2333/FloodGuard/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化接口包
func HandlerInit() {
	DB = database.DB
}
