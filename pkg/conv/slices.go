package conv

func ToIntSlice(ss []string) []int {
	var is []int
	for _, c := range ss {
		n := ToInt(c)
		if n == 0 {
			continue
		}
		is = append(is, n)
	}

	return is
}
