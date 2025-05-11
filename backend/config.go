package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	Port   string `env:"PORT" def:"8080"`
	JWTKey string `env:"JWT_KEY"`
	SSL    SSLConfig
	DB     DBConfig
	SMTP   SMTPConfig
}

type SSLConfig struct {
	Enable      bool   `env:"SSL_ENABLE" def:"false"`
	Certificate string `env:"SSL_CERTIFICATE"`
	Key         string `env:"SSL_KEY"`
}

type DBConfig struct {
	User     string `env:"DB_USER" def:"root"`
	Port     string `env:"DB_PORT" def:"3306"`
	Host     string `env:"DB_HOST" def:"127.0.0.1"`
	Name     string `env:"DB_NAME"`
	Password string `env:"DB_PASSWORD"`
}

type SMTPConfig struct {
	Server   string `env:"SMTP_SERVER"`
	Port     string `env:"SMTP_PORT" def:"25"`
	Mail     string `env:"SMTP_MAIL"`
	Password string `env:"SMTP_PASSWORD"`
}

func LoadConfigFromEnv(cfg interface{}) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 处理嵌套结构体
		if field.Kind() == reflect.Struct {
			if err := LoadConfigFromEnv(field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// 获取标签信息
		envKey := fieldType.Tag.Get("env")
		defVal := fieldType.Tag.Get("def")

		// 读取环境变量
		envVal := os.Getenv(envKey)
		if envVal == "" {
			envVal = defVal
		}

		// 跳过未设置且没有默认值的字段
		if envVal == "" {
			continue
		}

		// 设置字段值
		switch field.Kind() {
		case reflect.String:
			field.SetString(envVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(envVal)
			if err != nil {
				return fmt.Errorf("invalid bool value for %s: %v", envKey, err)
			}
			field.SetBool(boolVal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(envVal, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int value for %s: %v", envKey, err)
			}
			field.SetInt(intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(envVal, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid uint value for %s: %v", envKey, err)
			}
			field.SetUint(uintVal)
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(envVal, 64)
			if err != nil {
				return fmt.Errorf("invalid float value for %s: %v", envKey, err)
			}
			field.SetFloat(floatVal)
		default:
			return fmt.Errorf("unsupported type %s for field %s", field.Kind(), fieldType.Name)
		}
	}
	return nil
}
