package test

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTime_ToString(t *testing.T) {
	dt := dateTime.New(time.Date(2022, 9, 06, 21, 14, 25, 0, time.Local))
	assert.Equal(t, "2022-09-06 21:14:25", dt.ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "09-06 21:14:25", dt.ToString("MM-dd HH:mm:ss"))
	assert.Equal(t, "21:14:25", dt.ToString("HH:mm:ss"))

	assert.Equal(t, "2022/09/06 21:14:25", dt.ToString("yyyy/MM/dd HH:mm:ss"))
	assert.Equal(t, "09/06 21:14:25", dt.ToString("MM/dd HH:mm:ss"))

	assert.Equal(t, "09/06/2022 21:14:25", dt.ToString("MM/dd/yyyy HH:mm:ss"))

	assert.Equal(t, 2022, dt.Year())
	assert.Equal(t, 2022, dt.ToTime().Year())
	assert.Equal(t, 9, dt.Month())
	assert.Equal(t, time.Month(9), dt.ToTime().Month())
	assert.Equal(t, 06, dt.Day())
	assert.Equal(t, 06, dt.ToTime().Day())
	assert.Equal(t, 21, dt.Hour())
	assert.Equal(t, 21, dt.ToTime().Hour())
	assert.Equal(t, 14, dt.Minute())
	assert.Equal(t, 14, dt.ToTime().Minute())
	assert.Equal(t, 25, dt.Second())
	assert.Equal(t, 25, dt.ToTime().Second())

	assert.Equal(t, "2022-09-06 00:00:00", dt.Date().ToString("yyyy-MM-dd HH:mm:ss"))

	assert.Equal(t, "2024-09-06 21:14:25", dt.AddYears(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2021-09-06 21:14:25", dt.AddYears(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-11-06 21:14:25", dt.AddMonths(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-08-06 21:14:25", dt.AddMonths(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-08 21:14:25", dt.AddDays(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-05 21:14:25", dt.AddDays(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 23:14:25", dt.AddHours(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 20:14:25", dt.AddHours(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:16:25", dt.AddMinutes(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:13:25", dt.AddMinutes(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:14:27", dt.AddSeconds(2).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:14:24", dt.AddSeconds(-1).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:14:26", dt.AddMillisecond(1000).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 21:14:24", dt.AddMillisecond(-1000).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, float64(24), dt.Sub(dateTime.New(time.Date(2022, 9, 05, 21, 14, 25, 0, time.Local))).Hours())

	assert.Equal(t, "2023-11-09 21:14:25", dt.AddDate(1, 2, 3).ToString("yyyy-MM-dd HH:mm:ss"))
	assert.Equal(t, "2022-09-06 22:16:28", dt.AddTime(1, 2, 3).ToString("yyyy-MM-dd HH:mm:ss"))

	assert.Equal(t, time.Unix(100, 0).String(), dateTime.NewUnix(100).ToTime().String())
	assert.Equal(t, time.UnixMilli(100000).String(), dateTime.NewUnixMilli(100000).ToTime().String())
	dateTime.Since(dateTime.Now())
	dateTime.Now().Duration().String()
	assert.Equal(t, true, dateTime.Now().After(dt))
	assert.Equal(t, false, dateTime.Now().Before(dt))
	dateTime.Now().Value()
	now := dateTime.Now()
	now.Scan("2021-09-06 21:14:25")
	assert.Equal(t, "2021-09-06 21:14:25", now.ToString("yyyy-MM-dd HH:mm:ss"))
}
