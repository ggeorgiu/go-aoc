package conv

import "strconv"

func ToInt(val string) int {
	v, _ := strconv.Atoi(val)

	return v
}

func ToInt64(val string) int64 {
	v, _ := strconv.ParseInt(val, 10, 64)

	return v
}
