package util

import "time"

func GetCurrentMilliTime() int64 {
	return time.Now().UnixNano() / 1e6
}
