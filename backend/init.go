package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Loc *time.Location

// 初始化数据库
func Init() error {
	var err error

	if err := LoadConfigFromEnv(&Config); err != nil {
		return err
	}

	// 连接数据库
	if DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.Host,
		Config.DB.Port,
		Config.DB.Name,
	))); err != nil {
		return err
	}

	// 自动迁移数据库表
	for _, value := range Tables {
		DB.AutoMigrate(value)
	}
	Loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	GetMeta()

	return nil
}
