package gtools

import "strconv"

func I2a(v int64) string {
	return strconv.FormatInt(v, 10)
}

// uint to string base on 10
func Ui2a(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func B2a(v bool) string {
	return strconv.FormatBool(v)
}

func F2a(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}
