package dateTime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	time time.Time
}

// Now 当前时间
func Now() DateTime {
	return DateTime{
		time: time.Now(),
	}
}

// New 初始化
func New(time time.Time) DateTime {
	return DateTime{
		time: time,
	}
}

// New 初始化
func NewUnix(sec int64) DateTime {
	return DateTime{
		time: time.Unix(sec,0),
	}
}

// New 初始化
func NewUnixMilli(msec int64) DateTime {
	return DateTime{
		time: time.UnixMilli(msec),
	}
}

// Since time.Now().Sub(dt).
func Since(dt DateTime) time.Duration {
	return time.Since(dt.ToTime())
}

// ToString 转字符串，yyyy-MM-dd hh:mm:ss
func (receiver DateTime) ToString(format string) string {
	// 2006-01-02 15:04:05
	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "yy", "06", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "dd", "02", -1)

	format = strings.Replace(format, "hh", "15", -1)
	format = strings.Replace(format, "HH", "15", -1)
	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "ss", "05", -1)
	format = strings.Replace(format, "ffffff", "000000", -1)
	format = strings.Replace(format, "fff", "000", -1)

	return receiver.time.Format(format)
}

// Year 获取年
func (receiver DateTime) Year() int { return receiver.time.Year() }

// Month 获取月
func (receiver DateTime) Month() int { return int(receiver.time.Month()) }

// Day 获取日
func (receiver DateTime) Day() int { return receiver.time.Day() }

// Hour 获取小时
func (receiver DateTime) Hour() int { return receiver.time.Hour() }

// Minute 获取分钟
func (receiver DateTime) Minute() int { return receiver.time.Minute() }

// Second 获取秒
func (receiver DateTime) Second() int { return receiver.time.Second() }

// TotalSeconds 获取总秒数
func (receiver DateTime) TotalSeconds() float64 {
	return receiver.time.Sub(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)).Seconds()
}

// Duration 得到Duration
func (receiver DateTime) Duration() time.Duration {
	return receiver.time.Sub(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))
}

// UnixMilli 获取毫秒
func (receiver DateTime) UnixMilli() int64 { return receiver.time.UnixMilli() }

// UnixMicro 获取微秒
func (receiver DateTime) UnixMicro() int64 { return receiver.time.UnixMicro() }

// UnixNano 获取纳秒
func (receiver DateTime) UnixNano() int64 { return receiver.time.UnixNano() }

// Date 获取Date部份
func (receiver DateTime) Date() DateTime {
	year, month, day := receiver.time.Date()
	return New(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}

// AddDate 添加Date
func (receiver DateTime) AddDate(years int, months int, days int) DateTime {
	return New(receiver.time.AddDate(years, months, days))
}

// AddTime 添加Time
func (receiver DateTime) AddTime(hours int, minutes int, seconds int) DateTime {
	return New(receiver.time.Add(time.Hour * time.Duration(hours)).Add(time.Minute * time.Duration(minutes)).Add(time.Second * time.Duration(seconds)))
}

// AddYears 添加年
func (receiver DateTime) AddYears(year int) DateTime {
	return New(receiver.time.AddDate(year, 0, 0))
}

// AddMonths 添加月份
func (receiver DateTime) AddMonths(months int) DateTime {
	return New(receiver.time.AddDate(0, months, 0))
}

// AddDays 添加天数
func (receiver DateTime) AddDays(days int) DateTime {
	return New(receiver.time.AddDate(0, 0, days))
}

// AddHours 添加小时
func (receiver DateTime) AddHours(hours int) DateTime {
	return New(receiver.time.Add(time.Hour * time.Duration(hours)))
}

// AddMinutes 添加分钟
func (receiver DateTime) AddMinutes(minutes int) DateTime {
	return New(receiver.time.Add(time.Minute * time.Duration(minutes)))
}

// AddSeconds 添加秒
func (receiver DateTime) AddSeconds(seconds int) DateTime {
	return New(receiver.time.Add(time.Second * time.Duration(seconds)))
}

// AddMillisecond 添加毫秒
func (receiver DateTime) AddMillisecond(millisecond int) DateTime {
	return New(receiver.time.Add(time.Duration(millisecond) * time.Millisecond))
}

// Sub 时间相减
func (receiver DateTime) Sub(dt DateTime) time.Duration{
	return receiver.time.Sub(dt.time)
}

// ToTime 获取time.Time类型
func (receiver DateTime) ToTime() time.Time { return receiver.time }

// After 是否比dt时间大（晚）
func (receiver DateTime) After(dt DateTime) bool{
	return receiver.time.After(dt.time)
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.ToString("yyyy-MM-dd hh:mm:ss"))
}

// UnmarshalJSON to deserialize []byte
func (receiver *DateTime) UnmarshalJSON(b []byte) error {
	// 转time.Time
	layouts := []string{"2006-01-02 15:04:05", "2006-01-02", "2006-01-02T15:04:05Z07:00"}
	t := strings.Trim(string(b), "\"")
	for _, layout := range layouts {
		parse, err := time.ParseInLocation(layout, t, time.Local)
		if err == nil {
			receiver.time = parse
			return nil
		}
	}
	return fmt.Errorf("时间：%s 无法转换成time.Time类型", string(b))
}

// Value return json value, implement driver.Valuer interface
func (receiver DateTime) Value() (driver.Value, error) {
	return receiver.time, nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (receiver *DateTime) Scan(val any) error {
	if val == nil {
		return nil
	}
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}

	return receiver.UnmarshalJSON(ba)
}
