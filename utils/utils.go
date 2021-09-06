package utils

import (
	"strconv"
	"time"
)

type OftenTime time.Time

func (t OftenTime) MarshalJson() ([]byte, error) {
	now := time.Time(t)
	return []byte(now.Format("2006-01-02 15:04:05")), nil
}

// 时间相关
func TimestampNow() int64 {
	currentTime := time.Now()
	// currentTime.UnixNano()
	// microseconds 微秒
	// miliseconds 毫秒
	// unixNano 纳秒
	// 1秒=1000毫秒=1000,000微秒=1000,000,000纳秒
	// return currentTime.Unix() * 1000
	return currentTime.UnixNano() / 1e6 // 1e6=1000,000
}

// 字符串转数值
func Str2Int(s string) (int, error) {
	return strconv.Atoi(s)
}
func Str2Float(s string) (float64, error) {
	return strconv.ParseFloat(s, 10)
}
func Int2String(i int) string {
	return strconv.Itoa(i)
}
func Int642String(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

// 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}
func GetCurrentTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// http请求库：https://github.com/imroc/req
