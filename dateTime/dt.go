package dateTime

import (
	"database/sql/driver"

	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/farseer-go/fs/snc"
)

type DateTime struct {
	time time.Time
}

func Parse(layout, value string) (DateTime, error) {
	layout = strings.ReplaceAll(layout, "yyyy", "2006")
	layout = strings.ReplaceAll(layout, "yy", "06")
	layout = strings.ReplaceAll(layout, "MM", "01")
	layout = strings.ReplaceAll(layout, "dd", "02")
	layout = strings.ReplaceAll(layout, "HH", "15")
	layout = strings.ReplaceAll(layout, "mm", "04")
	layout = strings.ReplaceAll(layout, "ss", "05")
	layout = strings.ReplaceAll(layout, "ffffff", "000000")
	layout = strings.ReplaceAll(layout, "fff", "000")

	date, err := time.Parse(layout, value)
	if err != nil {
		return DateTime{}, err
	}
	return New(date), err
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

// NewUnix 初始化
func NewUnix(sec int64) DateTime {
	return DateTime{
		time: time.Unix(sec, 0),
	}
}

// NewUnixMilli 初始化
func NewUnixMilli(msec int64) DateTime {
	return DateTime{
		time: time.UnixMilli(msec),
	}
}

// NewUnixMicro 初始化
func NewUnixMicro(usec int64) DateTime {
	return DateTime{
		time: time.UnixMicro(usec),
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

//// TotalSeconds 获取总秒数
//func (receiver DateTime) TotalSeconds() float64 {
//	m, _ := time.ParseDuration(receiver.time.String())
//	return m.Seconds()
//	//return receiver.time.Sub(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)).Seconds()
//}

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

// AddSeconds 添加秒
func (receiver DateTime) Add(d time.Duration) DateTime {
	return New(receiver.time.Add(d))
}

// ResetDay 抹除天数，将天、小时、分、秒、毫秒 设为0
func (receiver DateTime) ResetDay() DateTime {
	return New(time.Date(receiver.Year(), time.Month(receiver.Month()), 1, 0, 0, 0, 0, time.Local))
}

// Sub 时间相减
func (receiver DateTime) Sub(dt DateTime) time.Duration {
	return receiver.time.Sub(dt.time)
}

// ToTime 获取time.Time类型
func (receiver DateTime) ToTime() time.Time { return receiver.time }

// After 是否比dt时间大（晚）
func (receiver DateTime) After(dt DateTime) bool {
	return receiver.time.After(dt.time)
}

// Before 是否比dt时间小（早）
func (receiver DateTime) Before(dt DateTime) bool {
	return receiver.time.Before(dt.time)
}

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (receiver DateTime) MarshalJSON() ([]byte, error) {
	return snc.Marshal(receiver.ToString("yyyy-MM-dd hh:mm:ss"))
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
	case time.Time:
		receiver.time = v
		return nil
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}

	return receiver.UnmarshalJSON(ba)
}
