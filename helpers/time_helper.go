package helpers

import "time"

// String2Time 解析字符串的日期时间为 time.Time 类型
func String2Time(str string) time.Time {
	pattern := "2006-01-02T15:04:05+08:00"
	t, err := time.Parse(pattern, str)

	if err != nil {
		return time.Unix(0, 1)
	}
	return t

}
