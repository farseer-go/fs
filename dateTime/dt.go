package dateTime

import (
	"encoding/json"
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

// ToString 转字符串，yyyy-MM-dd hh:mm:ss
func (d DateTime) ToString(format string) string {
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

	return d.time.Format(format)
}

// Year 获取年
func (d DateTime) Year() int { return d.time.Year() }

// Month 获取月
func (d DateTime) Month() int { return int(d.time.Month()) }

// Day 获取日
func (d DateTime) Day() int { return d.time.Day() }

// Hour 获取小时
func (d DateTime) Hour() int { return d.time.Hour() }

// Minute 获取分钟
func (d DateTime) Minute() int { return d.time.Minute() }

// Second 获取秒
func (d DateTime) Second() int { return d.time.Second() }

// UnixMilli 获取毫秒
func (d DateTime) UnixMilli() int64 { return d.time.UnixMilli() }

// UnixMicro 获取微秒
func (d DateTime) UnixMicro() int64 { return d.time.UnixMicro() }

// UnixNano 获取纳秒
func (d DateTime) UnixNano() int64 { return d.time.UnixNano() }

// Date 获取Date部份
func (d DateTime) Date() DateTime {
	year, month, day := d.time.Date()
	return New(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}

// AddDate 添加Date
func (d DateTime) AddDate(years int, months int, days int) DateTime {
	return New(d.time.AddDate(years, months, days))
}

// AddTime 添加Time
func (d DateTime) AddTime(hours int, minutes int, seconds int) DateTime {
	return New(d.time.Add(time.Hour * time.Duration(hours)).Add(time.Minute * time.Duration(minutes)).Add(time.Second * time.Duration(seconds)))
}

// AddYears 添加年
func (d DateTime) AddYears(year int) DateTime {
	return New(d.time.AddDate(year, 0, 0))
}

// AddMonths 添加月份
func (d DateTime) AddMonths(months int) DateTime {
	return New(d.time.AddDate(0, months, 0))
}

// AddDays 添加天数
func (d DateTime) AddDays(days int) DateTime {
	return New(d.time.AddDate(0, 0, days))
}

// AddHours 添加小时
func (d DateTime) AddHours(hours int) DateTime {
	return New(d.time.Add(time.Hour * time.Duration(hours)))
}

// AddMinutes 添加分钟
func (d DateTime) AddMinutes(minutes int) DateTime {
	return New(d.time.Add(time.Minute * time.Duration(minutes)))
}

// AddSeconds 添加秒
func (d DateTime) AddSeconds(seconds int) DateTime {
	return New(d.time.Add(time.Second * time.Duration(seconds)))
}

// ToTime 获取time.Time类型
func (d DateTime) ToTime() time.Time { return d.time }

// MarshalJSON to output non base64 encoded []byte
// 此处不能用指针，否则json序列化时不执行
func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ToString("yyyy-MM-dd hh:mm:ss"))
}

// UnmarshalJSON to deserialize []byte
func (d DateTime) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.time)
}
