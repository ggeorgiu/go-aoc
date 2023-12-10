package algs

func Sum(data []int) int {
	var sum int
	for _, v := range data {
		sum += v
	}

	return sum
}
