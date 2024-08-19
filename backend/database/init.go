package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func Init() error {
	var err error

	// 连接数据库
	if DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	))); err != nil {
		return err
	}

	// 自动迁移数据库表
	for _, value := range tables {
		DB.AutoMigrate(value)
	}

	// 为管理员用户授权
	if adminName := viper.GetString("ADMIN_NAME"); adminName != "" {
		var admin User
		if err := DB.First(
			&admin, "username = ?", adminName,
		).Error; err == nil {
			if !admin.Admin {
				DB.Model(&admin).Update("admin", true)
				fmt.Printf("已经成功将%v设置为管理员\n", adminName)
			}
		} else {
			fmt.Printf("无法为%v设置为管理员：%v\n", adminName, err.Error())
		}
	}
	return nil
}
