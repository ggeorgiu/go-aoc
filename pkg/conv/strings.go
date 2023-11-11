package conv

import "strconv"

func ToInt(val string) int {
	v, _ := strconv.Atoi(val)

	return v
}
