package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// 初始化配置文件
func Init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	for key, value := range defaultConfig {
		viper.SetDefault(key, value)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("无法找到配置文件，正在新建...")
			if err = viper.SafeWriteConfigAs("config.yaml"); err != nil {
				log.Fatalf("创建默认配置文件失败: %s\n", err)
			}
		} else {
			fmt.Println("无法解析配置文件，将重命名旧配置并新建...")
			if err := os.Rename(
				"config.yaml", "config.yaml.bak",
			); err != nil {
				log.Fatalf("重命名配置文件失败: %s\n", err)
			}
			if err = viper.SafeWriteConfigAs("config.yaml"); err != nil {
				log.Fatalf("创建新的配置文件失败: %s\n", err)
			}
		}
	} else {
		for key, value := range defaultConfig {
			if !viper.IsSet(key) {
				fmt.Printf("键 %s 不在配置文件里，正在新建\n", key)
				viper.Set(key, value)
			}
		}
		if err = viper.WriteConfig(); err != nil {
			log.Fatalf("写入配置文件失败: %s\n", err)
		}
	}
}
