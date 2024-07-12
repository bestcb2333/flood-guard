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
	new(Notice),
	new(RequestLog),
}

// 用户表
type User struct {
	gorm.Model
	Username      string
	Password      string
	Email         string
	Telephone     string
	Admin         bool
	NoticeWrited  []Notice     `gorm:"foreignKey:Writer;references:ID"`
	EventUploaded []FloodEvent `gorm:"foreignKey:Uploader;references:ID"`
}

// 区域表
type Region struct {
	gorm.Model
	Name        string
	FloodEvent  []FloodEvent  `gorm:"foreignKey:RegionID;references:ID"`
	HistoryData []HistoryData `gorm:"foreignKey:RegionID;references:ID"`
}

// 内涝事件
type FloodEvent struct {
	gorm.Model
	Region      Region `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint
	StartTime   time.Time
	EndTime     time.Time
	User        User `gorm:"foreignKey:Uploader;references:ID"`
	Uploader    uint `json:"-"`
	Severity    string
	Description string
}

// 历史数据
type HistoryData struct {
	ID         uint
	RecordTime time.Time
	Region     Region `gorm:"foreignKey:RegionID;references:ID"`
	RegionID   uint
	RainFall   float64
	WaterLevel float64
	Velocity   float64
}

// 通知公告
type Notice struct {
	gorm.Model
	User    User `gorm:"foreignKey:Writer;references:ID"`
	Author  uint `json:"-"`
	Title   string
	Content string
}

// 请求日志
type RequestLog struct {
	ID        uint
	CreatedAt time.Time
	ClientIP  string
	Path      string
	status    int
}
