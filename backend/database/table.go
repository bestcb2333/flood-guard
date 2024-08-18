package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
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

type CustomUint uint

func (cu CustomUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint(cu))
}

func (cu *CustomUint) UnmarshalJSON(data []byte) error {
	trimmedSlice := []byte(strings.Trim(string(data), `"`))
	var temp uint
	if err := json.Unmarshal(trimmedSlice, &temp); err != nil {
		return err
	}
	*cu = CustomUint(temp)
	return nil
}

func (cu CustomUint) Value() (driver.Value, error) {
	return uint(cu), nil
}

func (m *CustomUint) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		*m = CustomUint(v)
		return nil
	default:
		return fmt.Errorf("无法读取%T类型", value)
	}
}

type CustomBool bool

func (cb CustomBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(cb))
}

func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	trimmedSlice := []byte(strings.Trim(string(data), `"`))
	var temp bool
	if err := json.Unmarshal(trimmedSlice, &temp); err != nil {
		return err
	}
	*cb = CustomBool(temp)
	return nil
}

func (cb CustomBool) Value() (driver.Value, error) {
	return bool(cb), nil
}

func (cb *CustomBool) Scan(value any) error {
	switch v := value.(type) {
	case bool:
		*cb = CustomBool(v)
		return nil
	default:
		return fmt.Errorf("无法读取%T类型", value)
	}
}

type CustomFloat float64

func (cf CustomFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(cf))
}

func (cf *CustomFloat) UnmarshalJSON(data []byte) error {
	trimmedSlice := []byte(strings.Trim(string(data), `"`))
	var temp float64
	if err := json.Unmarshal(trimmedSlice, &temp); err != nil {
		return err
	}
	*cf = CustomFloat(temp)
	return nil
}

func (cf CustomFloat) Value() (driver.Value, error) {
	return float64(cf), nil
}

func (cf *CustomFloat) Scan(value any) error {
	switch v := value.(type) {
	case float64:
		*cf = CustomFloat(v)
		return nil
	default:
		return fmt.Errorf("无法读取%T类型", value)
	}
}

var timeFormat string = "2006-01-02 15:04:05"

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format(timeFormat))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsedTime, err := time.Parse(timeFormat, str)
	if err != nil {
		return fmt.Errorf("无法序列化时间%s: %v", str, err)
	}
	ct.Time = parsedTime
	return nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}
	return ct.Time.Format(time.RFC3339), nil
}

func (ct *CustomTime) Scan(value any) error {
	var t time.Time
	switch v := value.(type) {
	case time.Time:
		t = v
	case []byte:
		var err error
		t, err = time.Parse(time.RFC3339, string(v))
		if err != nil {
			return fmt.Errorf("无法将 %s 转换成time.Time类型", v)
		}
	case string:
		var err error
		t, err = time.Parse(time.RFC3339, v)
		if err != nil {
			return fmt.Errorf("无法将 %s 转换成time.Time类型", v)
		}
	default:
		return fmt.Errorf("CustomTime不支持的类型 %T", value)
	}
	ct.Time = t
	return nil
}

// 用户表
type User struct {
	ID            CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"comment:记录更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Username      string         `json:"username" gorm:"comment:用户名"`
	Password      string         `json:"password" gorm:"comment:密码"`
	Email         string         `json:"email" gorm:"comment:邮箱"`
	Profile       string         `json:"profile" gorm:"comment:个人简介"`
	Admin         CustomBool     `json:"admin" gorm:"comment:是管理员"`
	NoticeWritten []Notice       `json:"-" gorm:"foreignKey:Author;references:ID"`
	EventUploaded []FloodEvent   `json:"-" gorm:"foreignKey:Uploader;references:ID"`
}

// 区域表
type Region struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"comment:区域名"`
	Description string         `json:"description" gorm:"comment:区域描述"`
	Scope       string         `json:"scope" gorm:"comment:范围"`
	FloodEvent  []FloodEvent   `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	HistoryData []HistoryData  `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	SensorList  []Sensor       `json:"-" gorm:"foreignKey:RegionID;references:ID"`
}

// 内涝事件
type FloodEvent struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Region      Region         `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `json:"regionId" gorm:"comment:区域ID"`
	StartTime   CustomTime     `json:"startTime" gorm:"comment:开始时间"`
	EndTime     CustomTime     `json:"endTime" gorm:"comment:结束时间"`
	User        User           `json:"-" gorm:"foreignKey:Uploader;references:ID"`
	Uploader    CustomUint     `json:"uploader" gorm:"comment:上传者ID"`
	Severity    string         `json:"severity" gorm:"comment:严重性"`
	Position    string         `json:"position" gorm:"comment:位置"`
	Description string         `json:"description" gorm:"comment:描述"`
	Comment     []Comment      `json:"-" gorm:"foreignKey:Author;references:ID"`
}

// 历史数据
type HistoryData struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	RecordTime  CustomTime     `json:"recordTime" gorm:"comment:记录时间"`
	Region      Region         `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `json:"regionId" gorm:"comment:区域ID"`
	RainFall    CustomFloat    `json:"rainFall" gorm:"comment:降水量"`
	WaterLevel  CustomFloat    `json:"waterLevel" gorm:"comment:水位"`
	Velocity    CustomFloat    `json:"velocity" gorm:"comment:流速"`
	Temperature CustomFloat    `json:"temperature" gorm:"comment:气温"`
	Humidity    CustomFloat    `json:"humidity" gorm:"comment:湿度"`
	DataSource  string         `json:"dataSource" gorm:"comment:数据源"`
}

// 通知公告
type Notice struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        User           `json:"-" gorm:"foreignKey:Author;references:ID"`
	Author      CustomUint     `json:"-" gorm:"comment:作者ID"`
	Title       string         `json:"title" gorm:"comment:标题"`
	Content     string         `json:"content" gorm:"comment:内容"`
	Importrance CustomUint     `json:"importance" gorm:"comment:重要性"`
	Comment     []Comment      `json:"comment" gorm:"foreignKey:Author;references:ID"`
}

// 用户评论
type Comment struct {
	ID        CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:创建时间"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User           `json:"-" gorm:"foreignKey:Author;references:ID"`
	Author    CustomUint     `json:"Author" gorm:"comment:作者ID"`
	Content   string         `json:"content" gorm:"comment:内容"`
	Related   string         `json:"related" gorm:"comment:关联的表"`
}

// 传感器
type Sensor struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"comment:名称"`
	Abscissa    CustomFloat    `json:"abscissa" gorm:"comment:横坐标"`
	Ordinate    CustomFloat    `json:"ordinate" gorm:"comment:纵坐标"`
	Description string         `json:"description" gorm:"comment:描述"`
	Region      Region         `json:"-" gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `json:"regionId" gorm:"comment:区域ID"`
}

// 传感器状态
type SensorStatus struct {
	ID          CustomUint     `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Time        CustomTime     `json:"time" gorm:"comment:时间"`
	Sensor      Sensor         `json:"-" gorm:"foreignKey:SensorID;references:ID"`
	SensorID    CustomUint     `json:"sensorId" gorm:"comment:传感器ID"`
	Status      string         `json:"status" gorm:"comment:状态"`
	Description string         `json:"description" gorm:"comment:描述"`
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
