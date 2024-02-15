package utils

import "time"

const (
	TimeYMDHIS   = "2006-01-02 15:04:05"
	TimeYMD      = "2006-01-02"
	TimeHIS      = "15:04:05"
	AsiaShanghai = "Asia/Shanghai"

	TimeFormat       = "15:04:05"
	DateFormat       = "2006-01-02"
	DateTimeFormat   = "2006-01-02 15:04:05"
	DateTimeTFormat  = "2006-01-02T15:04:05"
	DateTimeFormatEn = "2006-01-02 at 15:04:05"
)

// 获取时间范伟 Time类型 2020-08-18 00:00:00 to time.Time
func TimeStrToTimestamp(timeStr string) (beginTime time.Time) {
	beginTime, _ = time.ParseInLocation(TimeYMDHIS, timeStr, time.Local)
	return beginTime
}
