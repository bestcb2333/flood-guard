package config

import (
	"os"

	"github.com/spf13/viper"
)

// 初始化配置文件
func Init() {
	viper.AutomaticEnv()
	for key, value := range Config {
		if temp, ok := os.LookupEnv(key); ok {
			viper.Set(key, temp)
		} else {
			viper.Set(key, value)
		}
	}
}
