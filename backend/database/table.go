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
	Username      string       `json:"username" gorm:"comment:用户名"`
	Password      string       `json:"password" gorm:"comment:密码"`
	Email         string       `json:"email" gorm:"comment:邮箱"`
	Profile       string       `json:"profile" gorm:"comment:个人简介"`
	Admin         bool         `json:"admin" gorm:"comment:是管理员"`
	NoticeWritten []Notice     `json:"-" gorm:"foreignKey:Author;references:ID"`
	EventUploaded []FloodEvent `json:"-" gorm:"foreignKey:Uploader;references:ID"`
}

// 区域表
type Region struct {
	gorm.Model
	Name        string        `json:"name" gorm:"comment:区域名"`
	Description string        `json:"description" gorm:"comment:区域描述"`
	Scope       string        `json:"scope" gorm:"comment:范围"`
	FloodEvent  []FloodEvent  `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	HistoryData []HistoryData `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	SensorList  []Sensor      `json:"-" gorm:"foreignKey:RegionID;references:ID"`
}

// 内涝事件
type FloodEvent struct {
	gorm.Model
	Region      Region    `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint      `json:"regionId" gorm:"comment:区域ID"`
	StartTime   time.Time `json:"startTime" gorm:"comment:开始时间"`
	EndTime     time.Time `json:"endTime" gorm:"comment:结束时间"`
	User        User      `json:"-" gorm:"foreignKey:Uploader;references:ID"`
	Uploader    uint      `json:"uploader" gorm:"comment:上传者ID"`
	Severity    string    `json:"severity" gorm:"comment:严重性"`
	Position    string    `json:"position" gorm:"comment:位置"`
	Description string    `json:"description" gorm:"comment:描述"`
	Comment     []Comment `json:"-" gorm:"foreignKey:Author;references:ID"`
}

// 历史数据
type HistoryData struct {
	ID          uint
	RecordTime  time.Time `json:"recordTime" gorm:"comment:记录时间"`
	Region      Region    `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint      `json:"regionId" gorm:"comment:区域ID"`
	RainFall    float64   `json:"rainFall" gorm:"comment:降水量"`
	WaterLevel  float64   `json:"waterLevel" gorm:"comment:水位"`
	Velocity    float64   `json:"velocity" gorm:"comment:流速"`
	Temperature float64   `json:"temperature" gorm:"comment:气温"`
	Humidity    float64   `json:"humidity" gorm:"comment:湿度"`
	DataSource  string    `json:"dataSource" gorm:"comment:数据源"`
}

// 通知公告
type Notice struct {
	gorm.Model
	User        User      `json:"-" gorm:"foreignKey:Author;references:ID"`
	Author      uint      `json:"-" gorm:"comment:作者ID"`
	Title       string    `json:"title" gorm:"comment:标题"`
	Content     string    `json:"content" gorm:"comment:内容"`
	Importrance int       `json:"importance" gorm:"comment:重要性"`
	Comment     []Comment `json:"comment" gorm:"foreignKey:Author;references:ID"`
}

// 用户评论
type Comment struct {
	gorm.Model
	User    User   `json:"-" gorm:"foreignKey:Author;references:ID"`
	Author  uint   `json:"Author" gorm:"comment:作者ID"`
	Content string `json:"content" gorm:"comment:内容"`
	Related string `json:"related" gorm:"comment:关联的表"`
}

// 传感器
type Sensor struct {
	gorm.Model
	Name        string  `json:"name" gorm:"comment:名称"`
	Abscissa    float64 `json:"abscissa" gorm:"comment:横坐标"`
	Ordinate    float64 `json:"ordinate" gorm:"comment:纵坐标"`
	Description string  `json:"description" gorm:"comment:描述"`
	Region      Region  `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    uint    `json:"regionId" gorm:"comment:区域ID"`
}

// 传感器状态
type SensorStatus struct {
	ID          uint
	Time        time.Time `json:"time" gorm:"comment:时间"`
	Sensor      Sensor    `json:"-" gorm:"foreignKey:SensorID;references:ID"`
	SensorID    uint      `json:"sensorId" gorm:"comment:传感器ID"`
	Status      string    `json:"status" gorm:"comment:状态"`
	Description string    `json:"description" gorm:"comment:描述"`
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
