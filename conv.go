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

func A2I(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func A2Ui(v string) (uint64, error) {
	return strconv.ParseUint(v, 10, 64)
}

func A2F(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func A2B(v string) (bool, error) {
	return strconv.ParseBool(v)
}
