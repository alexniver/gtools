package util

import "time"

// NowInMillisecond return time now in Millisecond
func NowInMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
