package database

import (
	"time"

	"gorm.io/gorm"
)

var tables = []any{
	new(User),
	new(Region),
	new(FloodEvent),
	new(HistoryData),
	new(RequestLog),
}

// 用户表
type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	Telephone string
	Admin     bool
}

// 区域表
type Region struct {
	gorm.Model
	Name        string
	FloodEvent  []FloodEvent
	HistoryData []HistoryData
}

// 内涝事件
type FloodEvent struct {
	gorm.Model
	Region
	StartTime   time.Time
	EndTime     time.Time
	Severity    string
	Description string
}

// 历史数据
type HistoryData struct {
	ID         uint
	RecordTime time.Time
	Region
	RainFall   float64
	WaterLevel float64
}

// 请求日志
type RequestLog struct {
	ID          uint
	CreatedAt   time.Time
	ClientIP    string
	Path        string
	status      int
	RequestBody string
}
