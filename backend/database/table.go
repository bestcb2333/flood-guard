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
	case int64:
		if v == 0 {
			*cb = CustomBool(false)
		} else {
			*cb = CustomBool(true)
		}
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
	ID            CustomUint     `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"comment:创建时间"`
	UpdatedAt     time.Time      `gorm:"comment:记录更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Username      string         `gorm:"comment:用户名"`
	Password      string         `gorm:"comment:密码"`
	Email         string         `gorm:"comment:邮箱"`
	Profile       string         `gorm:"comment:个人简介"`
	Admin         CustomBool     `gorm:"comment:是管理员"`
	NoticeWritten []Notice       `gorm:"foreignKey:Author;references:ID"`
	EventUploaded []FloodEvent   `gorm:"foreignKey:Uploader;references:ID"`
	Thread        []Thread       `gorm:"foreignKey:ThreadID;references:ID"`
}

// 区域表
type Region struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"comment:区域名"`
	Description string         `gorm:"comment:区域描述"`
	Scope       string         `gorm:"comment:范围"`
	FloodEvent  []FloodEvent   `gorm:"foreignKey:RegionID;references:ID"`
	HistoryData []HistoryData  `gorm:"foreignKey:RegionID;references:ID"`
	SensorList  []Sensor       `gorm:"foreignKey:RegionID;references:ID"`
}

// 内涝事件
type FloodEvent struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Region      Region         `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `gorm:"comment:区域ID"`
	StartTime   CustomTime     `gorm:"comment:开始时间"`
	EndTime     CustomTime     `gorm:"comment:结束时间"`
	User        User           `gorm:"foreignKey:Uploader;references:ID"`
	Uploader    CustomUint     `gorm:"comment:上传者ID"`
	Severity    string         `gorm:"comment:严重性"`
	Position    string         `gorm:"comment:位置"`
	Description string         `gorm:"comment:描述"`
	Comment     []Comment      `gorm:"foreignKey:Author;references:ID"`
}

// 历史数据
type HistoryData struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Region      Region         `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `gorm:"comment:区域ID"`
	RainFall    CustomFloat    `gorm:"comment:降水量"`
	WaterLevel  CustomFloat    `gorm:"comment:水位"`
	Velocity    CustomFloat    `gorm:"comment:流速"`
	Temperature CustomFloat    `gorm:"comment:气温"`
	Humidity    CustomFloat    `gorm:"comment:湿度"`
	DataSource  string         `gorm:"comment:数据源"`
}

// 通知公告
type Notice struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        User           `gorm:"foreignKey:Author;references:ID"`
	Author      CustomUint     `gorm:"comment:作者ID"`
	Title       string         `gorm:"comment:标题"`
	Content     string         `gorm:"comment:内容"`
	Importrance CustomUint     `gorm:"comment:重要性"`
	Comment     []Comment      `gorm:"foreignKey:Author;references:ID"`
}

// 用户评论
type Comment struct {
	ID        CustomUint     `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User           `gorm:"foreignKey:Author;references:ID"`
	Author    CustomUint     `gorm:"comment:作者ID"`
	Content   string         `gorm:"comment:内容"`
	Related   string         `gorm:"comment:关联的表"`
}

// 传感器
type Sensor struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"comment:记录更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"comment:名称"`
	Abscissa    CustomFloat    `gorm:"comment:横坐标"`
	Ordinate    CustomFloat    `gorm:"comment:纵坐标"`
	Description string         `gorm:"comment:描述"`
	Region      Region         `gorm:"foreignKey:RegionID;references:ID"`
	RegionID    CustomUint     `gorm:"comment:区域ID"`
}

// 传感器状态
type SensorStatus struct {
	ID          CustomUint     `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Time        CustomTime     `gorm:"comment:时间"`
	Sensor      Sensor         `gorm:"foreignKey:SensorID;references:ID"`
	SensorID    CustomUint     `gorm:"comment:传感器ID"`
	Status      string         `gorm:"comment:状态"`
	Description string         `gorm:"comment:描述"`
}

type Thread struct {
	ID         CustomUint     `gorm:"primaryKey"`
	CreatedAt  time.Time      `gorm:"comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"comment:更新时间"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ThreadID   string         `gorm:"comment:会话ID"`
	ThreadName string         `gorm:"会话名"`
	User       User           `gorm:"foreignKey:UserID;references:ID"`
	UserID     CustomUint     `gorm:"comment:用户ID"`
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
