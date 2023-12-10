package conv

func ToIntSlice(ss []string) []int {
	var is []int
	for _, c := range ss {
		if c == "" {
			continue
		}

		n := ToInt(c)
		is = append(is, n)
	}

	return is
}
