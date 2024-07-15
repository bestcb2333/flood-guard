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
	new(Comment),
	new(Sensor),
	new(SensorStatus),
	new(RequestLog),
}

// 用户表
type User struct {
	gorm.Model
	Username      string
	Password      string
	Email         string
	Profile       string
	Admin         bool
	NoticeWrited  []Notice     `gorm:"foreignKey:Writer;references:ID"`
	EventUploaded []FloodEvent `gorm:"foreignKey:Uploader;references:ID"`
}

// 区域表
type Region struct {
	gorm.Model
	Name        string
	Description string
	Scope       string
	FloodEvent  []FloodEvent  `gorm:"foreignKey:RegionID;references:ID"`
	HistoryData []HistoryData `gorm:"foreignKey:RegionID;references:ID"`
	SensorList  []Sensor      `gorm:"foreignKey:RegionID;references:ID"`
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
	Position    string
	Description string
	Comment     []Comment `gorm:"foreignKey:Author;references:ID"`
}

// 历史数据
type HistoryData struct {
	ID          uint
	RecordTime  time.Time
	Region      Region `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint
	RainFall    float64
	WaterLevel  float64
	Velocity    float64
	Temperature float64
	Humidity    float64
	DataSource  string
}

// 通知公告
type Notice struct {
	gorm.Model
	User        User `gorm:"foreignKey:Author;references:ID"`
	Author      uint `json:"-"`
	Title       string
	Content     string
	Importrance int
	Comment     []Comment `gorm:"foreignKey:Author;references:ID"`
}

// 用户评论
type Comment struct {
	gorm.Model
	User    User `gorm:"foreignKey:Author;references:ID"`
	Author  uint `json:"-"`
	Content string
	Related string
}

// 传感器
type Sensor struct {
	gorm.Model
	Name        string
	Abscissa    float64
	Ordinate    float64
	Description string
	Region      Region `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint
}

// 传感器状态
type SensorStatus struct {
	ID          uint
	Time        time.Time
	Sensor      Sensor `gorm:"foreignKey:SensorID;references:ID"`
	SensorID    uint
	status      string
	Description string
}

// 请求日志
type RequestLog struct {
	ID        uint
	CreatedAt time.Time
	ClientIP  string
	Path      string
	status    int
	Method    string
	UserID    uint
}
