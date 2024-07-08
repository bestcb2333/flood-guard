package database

import (
	"time"

	"gorm.io/gorm"
)

var tables = []any{
	new(User), new(FloodEvent), new(HistoryData), new(AlertRecord),
}

// 用户表
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Admin    bool
}

// 内涝事件
type FloodEvent struct {
	gorm.Model
	Location    string
	StartTime   time.Time
	EndTime     time.Time
	Severity    string
	Description string
}

// 历史数据
type HistoryData struct {
	gorm.Model
	Location   string
	RainFall   float64
	WaterLevel float64
	Date       time.Time `gorm:"type:date"`
}

// 预警记录
type AlertRecord struct {
	gorm.Model
	AlertTime  time.Time
	AlertLevel string
	Message    string
}

type RequestLog struct {
	ID          uint
	CreatedAt   time.Time
	ClientIP    string
	Path        string
	status      int
	RequestBody string
}
