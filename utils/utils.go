package utils

import "time"

type OftenTime time.Time

func (t OftenTime) MarshalJson() ([]byte, error) {
	now := time.Time(t)
	return []byte(now.Format("2006-01-02 15:04:05")), nil
}

// 时间相关
func TimestampNow() int64 {
	currentTime := time.Now()
	currentTime.UnixMicro()
	// microseconds 微秒
	// miliseconds 毫秒
	// unixNano 纳秒
	// 1秒=1000毫秒=1000,000微秒=1000,000,000纳秒
	// return currentTime.Unix() * 1000
	return currentTime.UnixNano() / 1e6 // 1e6=1000,000
}
