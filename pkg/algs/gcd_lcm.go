package algs

// GCD Calculates greatest common divisor using Euclid's algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// LCM Calculates least common multiple for an array of integers
func LCM(integers ...int) int {
	res := integers[0]

	for i := 1; i < len(integers); i++ {
		res = res * integers[i] / GCD(res, integers[i])
	}

	return res
}
