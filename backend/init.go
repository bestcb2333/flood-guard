package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func InitDB(config *Config) (*gorm.DB, error) {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)))
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库表
	for _, value := range Tables {
		db.AutoMigrate(value)
	}

	return db, nil
}
