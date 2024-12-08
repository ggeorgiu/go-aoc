package gen

func CartesianProduct(seed string, current string, desiredLen int, result *[]string) {
	if len(current) == desiredLen {
		*result = append(*result, current)
		return
	}

	for _, c := range seed {
		CartesianProduct(seed, current+string(c), desiredLen, result)
	}
}
