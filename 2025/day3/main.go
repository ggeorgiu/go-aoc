package main

import (
	"fmt"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2025/day3/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	var sum int
	for _, line := range input {
		var mx int32
		var local int32

		for i, c1 := range line {
			for _, c2 := range line[i+1:] {
				d1 := c1 - '0'
				d2 := c2 - '0'

				local = d1*10 + d2
				if local > mx {
					mx = local
				}
			}
		}

		sum += int(mx)
	}

	return sum
}

func part2(input []string) any {
	var total int
	for _, line := range input {
		var nr int
		digits := make([]int32, 0, len(line))

		for _, c := range line {
			digits = append(digits, c-'0')
		}

		limit := 12
		maxPos := 0
		for limit > 0 {
			dig, pos := maxIntervalDigit(digits, maxPos, len(digits)-limit)
			nr = (nr * 10) + int(dig)
			maxPos = pos + 1
			limit--
		}
		total += nr
	}

	return total
}

func maxIntervalDigit(digits []int32, from, to int) (int32, int) {
	var md int32
	var idx int
	for i := from; i <= to; i++ {
		if digits[i] > md {
			md = digits[i]
			idx = i
		}
	}

	return md, idx
}
