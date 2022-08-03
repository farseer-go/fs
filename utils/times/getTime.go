package times

import "time"

// GetTime 根据time.Duration转换成天、小时、分钟、秒
func GetTime(d time.Duration) (days int, hours int, minutes int, seconds int) {
	seconds = int(d.Milliseconds() / 1000)
	if seconds > -60 {
		minutes = seconds / 60
		seconds -= minutes * 60
	}

	if minutes > -60 {
		hours = minutes / 60
		minutes -= hours * 60
	}

	if hours > -24 {
		days = hours / 24
		hours -= days * 24
	}

	return days, hours, minutes, seconds
}

// GetDate 获取当前日期
func GetDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}
