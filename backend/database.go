package main

import (
	"time"

	"github.com/paulmach/orb"
	"gorm.io/gorm"
)

var Tables = []any{
	new(User),
	new(Region),
	new(Event),
	new(History),
	new(Resource),
	new(Notice),
	new(Sensor),
	new(Analysis),
}

// 用户表
type User struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
	Name      string         `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:用户名"`
	Password  string         `json:"-" gorm:"type:VARCHAR(255);not null;comment:密码"`
	Email     string         `json:"email" gorm:"type:VARCHAR(255);not null;unique;comment:邮箱"`
	Avatar    *string        `json:"avatar" gorm:"type:VARCHAR(255);unique;comment:头像路径"`
	Profile   *string        `json:"profile" gorm:"type:TEXT;comment:个人简介"`
	Admin     bool           `json:"admin" gorm:"not null;comment:是管理员"`
	RegionID  *uint          `json:"-" gorm:"index;comment:志愿服务的区域"`
	Region    *Region        `json:"region" extend:"true"`
	Notices   []Notice       `json:"notices" gorm:"constraint:OnDelete:SET NULL"`
	Events    []Event        `json:"events" gorm:"constraint:OnDelete:SET NULL"`
	Analyses  []Analysis     `json:"analyses" gorm:"constraint:OnDelete:CASCADE"`
}

// 区域表
type Region struct {
	ID          uint             `json:"id" gorm:"primarykey;comment:ID"`
	UpdatedAt   time.Time        `json:"updatedAt" gorm:"comment:更新时间"`
	Name        string           `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:区域名"`
	Description *string          `json:"description" gorm:"type:TEXT;comment:区域描述"`
	Altitude    *uint            `json:"altitude" gorm:"comment:海拔"`
	Drainage    *uint            `json:"drainage" gorm:"comment:排水系统完善度"`
	Forecast    *uint            `json:"forecast" gorm:"comment:预报雨量"`
	Coordinate  orb.MultiPolygon `json:"coordinate" gorm:"type:JSON;serializer:json;comment:坐标范围"`
	Users       []User           `json:"users" gorm:"constraint:OnDelete:SET NULL"`
	Events      []Event          `json:"events" gorm:"many2many:region_events;constraint:OnDelete:CASCADE"`
	Histories   []History        `json:"histories" gorm:"constraint:OnDelete:CASCADE"`
	Resources   []Resource       `json:"resources" gorm:"constraint:OnDelete:CASCADE"`
	Sensors     []Sensor         `json:"sensors" gorm:"constraint:OnDelete:CASCADE"`
}

// 内涝事件
type Event struct {
	ID          uint             `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt   time.Time        `json:"createdAt" gorm:"comment:创建时间"`
	Regions     []Region         `json:"regions" gorm:"many2many:region_events;constraint:OnDelete:CASCADE"`
	StartTime   time.Time        `json:"startTime" gorm:"not null;comment:开始时间"`
	EndTime     time.Time        `json:"endTime" gorm:"not null;comment:结束时间"`
	User        *User            `json:"user" extend:"true"`
	UserID      *uint            `json:"-" gorm:"index;comment:上传的用户ID"`
	Severity    uint             `json:"severity" gorm:"not null;comment:严重性"`
	Coordinate  orb.MultiPolygon `json:"coordinate" gorm:"type:JSON;serializer:json;comment:坐标"`
	Description *string          `json:"description" gorm:"type:TEXT;comment:描述"`
}

// 历史数据
type History struct {
	ID          uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	RegionID    uint      `json:"-" gorm:"not null;index;comment:相关的区域ID"`
	Region      Region    `json:"region" extend:"true"`
	Rainfall    *float64  `json:"rainfall" gorm:"comment:降水量"`
	Waterlevel  *float64  `json:"waterlevel" gorm:"comment:水位"`
	Description *string   `json:"description" gorm:"type:TEXT;comment:描述"`
}

// 资源
type Resource struct {
	ID         uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt  time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Type       string    `json:"type" gorm:"type:VARCHAR(255);not null;comment:资源类型"`
	Name       string    `json:"name" gorm:"type:VARCHAR(255);not null;comment:资源名称"`
	Quantity   uint      `json:"quantity" gorm:"not null;comment:数量"`
	RegionID   uint      `json:"-" gorm:"not null;index;comment:所在的区域ID"`
	Region     Region    `json:"region" extend:"true"`
	Coordinate orb.Point `json:"coordinate" gorm:"type:JSON;serializer:json;comment:坐标"`
	Available  bool      `json:"available" gorm:"not null;comment:是否可用"`
}

// 通知公告
type Notice struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
	User      *User     `json:"user" extend:"true"`
	UserID    *uint     `json:"-" gorm:"index;comment:编写者ID"`
	Title     string    `json:"title" gorm:"type:VARCHAR(255);not null;comment:标题"`
	Content   *string   `json:"content" gorm:"type:TEXT;comment:内容"`
}

// 传感器
type Sensor struct {
	ID          uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:创建时间"`
	Name        string    `json:"name" gorm:"type:VARCHAR(255);not null;unique;comment:名称"`
	Coordinate  orb.Point `json:"coordinate" gorm:"type:JSON;serializer:json;comment:坐标"`
	Description *string   `json:"description" gorm:"type:TEXT;comment:描述"`
	Available   bool      `json:"available" gorm:"not null;comment:是否可用"`
	RegionID    uint      `json:"-" gorm:"not null;index;comment:所在的区域ID"`
	Region      Region    `json:"region" extend:"true"`
}

// LLM智能分析记录
type Analysis struct {
	ID        uint      `json:"id" gorm:"primarykey;comment:ID"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UserID    uint      `json:"-" gorm:"not null;index;comment:发起的用户ID"`
	User      User      `json:"user" extend:"true"`
	Role      string    `json:"role" gorm:"type:VARCHAR(255);not null;comment:消息类型"`
	Content   *string   `json:"content" gorm:"type:TEXT;comment:内容"`
}
