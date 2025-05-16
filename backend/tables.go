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
	new(Message),
}

type ListDTO struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func (ld ListDTO) GetListDTO() *ListDTO {
	return &ld
}

type DeleteDTO struct {
	IDs []uint `form:"id"`
}

type IDField struct {
	ID uint `json:"id" gorm:"primarykey;comment:ID"`
}

type CreatedAtField struct {
	CreatedAt time.Time `json:"createdAt" gorm:"not null;comment:创建时间"`
}

type UpdatedAtField struct {
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;comment:更新时间"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index;comment:删除时间"`
}

// 用户表
type User struct {
	IDField
	CreatedAtField
	UpdatedAtField
	Name     string    `json:"name" gorm:"type:VARCHAR(20);not null;unique;comment:用户名"`
	Password string    `json:"-" gorm:"type:CHAR(64);not null;comment:密码"`
	Email    string    `json:"email" gorm:"type:VARCHAR(50);not null;unique;comment:邮箱"`
	Avatar   string    `json:"avatar" gorm:"type:VARCHAR(50);not null;comment:头像路径"`
	Profile  string    `json:"profile" gorm:"type:VARCHAR(200);not null;comment:个人简介"`
	Admin    bool      `json:"admin" gorm:"not null;comment:是管理员"`
	RegionID *uint     `json:"regionId" gorm:"index;comment:志愿服务的区域"`
	Region   *Region   `json:"region"`
	Notices  []Notice  `json:"notices" gorm:"constraint:OnDelete:SET NULL"`
	Events   []Event   `json:"events" gorm:"constraint:OnDelete:SET NULL"`
	Messages []Message `json:"messages" gorm:"constraint:OnDelete:CASCADE"`
}

// 区域表
type RegionDTO struct {
	Name        string           `json:"name" gorm:"type:VARCHAR(20);not null;unique;comment:区域名"`
	Description string           `json:"description,omitempty" gorm:"type:VARCHAR(200);not null;comment:区域描述"`
	Altitude    *uint            `json:"altitude,omitempty" gorm:"comment:海拔"`
	Drainage    *uint            `json:"drainage,omitempty" gorm:"comment:排水系统完善度"`
	Forecast    *uint            `json:"forecast,omitempty" gorm:"comment:预报雨量"`
	Coordinate  orb.MultiPolygon `json:"coordinate,omitempty" gorm:"type:JSON;serializer:json;not null;comment:坐标范围"`
}

type Region struct {
	IDField
	UpdatedAtField
	RegionDTO
	Users     []User     `json:"users,omitempty" gorm:"constraint:OnDelete:SET NULL"`
	Events    []Event    `json:"events,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Histories []History  `json:"histories,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Resources []Resource `json:"resources,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	Sensors   []Sensor   `json:"sensors,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

// 内涝事件
type EventDTO struct {
	Name        string     `json:"name" gorm:"type:VARCHAR(20);not null;unique;comment:事件名称"`
	RegionID    uint       `json:"regionId" gorm:"not null;index;comment:所在区域ID"`
	StartTime   time.Time  `json:"startTime" gorm:"not null;comment:开始时间"`
	EndTime     *time.Time `json:"endTime" gorm:"comment:结束时间"`
	Severity    string     `json:"severity" gorm:"type:VARCHAR(20);not null;comment:严重性"`
	Coordinate  orb.Point  `json:"coordinate" gorm:"type:JSON;serializer:json;not null;comment:坐标"`
	Description string     `json:"description" gorm:"type:TEXT;not null;comment:描述"`
}

type Event struct {
	IDField
	CreatedAtField
	*EventDTO
	Region *Region `json:"region" gorm:"constraint:OnDelete:CASCADE"`
	UserID *uint   `json:"userId" gorm:"index;comment:上传的用户ID"`
	User   *User   `json:"user"`
}

// 历史数据
type HistoryDTO struct {
	RegionID    uint     `json:"regionId" gorm:"not null;index;comment:相关的区域ID"`
	Rainfall    *float64 `json:"rainfall,omitempty" gorm:"comment:降水量"`
	Waterlevel  *float64 `json:"waterlevel,omitempty" gorm:"comment:水位"`
	Source      string   `json:"source,omitempty" gorm:"type:VARCHAR(20);not null;comment:数据源"`
	Description string   `json:"description,omitempty" gorm:"type:VARCHAR(200);not null;comment:描述"`
}

type History struct {
	IDField
	CreatedAtField
	HistoryDTO
	Region *Region `json:"region,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

// 资源
type ResourceDTO struct {
	Type       string    `json:"type" gorm:"type:VARCHAR(20);not null;comment:资源类型"`
	Name       string    `json:"name" gorm:"type:VARCHAR(50);not null;comment:资源名称"`
	Quantity   uint      `json:"quantity" gorm:"not null;comment:数量"`
	RegionID   uint      `json:"regionId" gorm:"not null;index;comment:所在的区域ID"`
	Coordinate orb.Point `json:"coordinate" gorm:"type:JSON;serializer:json;not null;comment:坐标"`
	Available  bool      `json:"available" gorm:"not null;comment:是否可用"`
}

type Resource struct {
	IDField
	UpdatedAtField
	ResourceDTO
	Region *Region `json:"region" gorm:"constraint:OnDelete:CASCADE"`
}

// 通知公告
type NoticeDTO struct {
	Title   string `json:"title" gorm:"type:VARCHAR(20);not null;comment:标题"`
	Content string `json:"content" gorm:"type:VARCHAR(200);not null;comment:内容"`
}

type Notice struct {
	IDField
	CreatedAtField
	UpdatedAtField
	*NoticeDTO
	UserID *uint `json:"userId" gorm:"index;comment:编写者ID"`
	User   *User `json:"user"`
}

// 传感器
type SensorDTO struct {
	Name        string    `json:"name" gorm:"type:VARCHAR(20);not null;unique;comment:名称"`
	Coordinate  orb.Point `json:"coordinate" gorm:"type:JSON;serializer:json;not null;comment:坐标"`
	Description string    `json:"description" gorm:"type:VARCHAR(200);not null;comment:描述"`
	Available   bool      `json:"available" gorm:"not null;comment:是否可用"`
	RegionID    uint      `json:"regionId" gorm:"not null;index;comment:所在的区域ID"`
}

type Sensor struct {
	IDField
	CreatedAtField
	*SensorDTO
	Region Region `json:"region"`
}

// LLM智能分析记录
type MessageDTO struct {
	Role    string `json:"role" gorm:"type:VARCHAR(20);not null;comment:消息类型"`
	Content string `json:"content" gorm:"type:TEXT;not null;comment:内容"`
}

type Message struct {
	IDField
	CreatedAtField
	MessageDTO
	UserID uint `json:"regionId" gorm:"not null;index;comment:发起的用户ID"`
	User   User `json:"user"`
}
