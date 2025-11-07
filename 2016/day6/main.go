package main

import (
	"fmt"
	"math"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day6/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	size := len(input[0])

	result := make([]uint8, size)
	for i := 0; i < size; i++ {
		count := make([]int, 26)
		for _, line := range input {
			c := line[i]
			count[c-'a']++
		}

		var m int
		var mr uint8
		for i, v := range count {
			if v > m {
				m = v
				mr = uint8(i + 'a')
			}
		}

		for i, v := range count {
			fmt.Printf("%s:%d ", string(rune(i+'a')), v)
		}
		fmt.Println()

		result[i] = mr
	}

	return string(result)
}

func part2(input []string) any {
	size := len(input[0])

	result := make([]uint8, size)
	for i := 0; i < size; i++ {
		count := make([]int, 26)
		for _, line := range input {
			c := line[i]
			count[c-'a']++
		}

		var m = math.MaxInt
		var mr uint8
		for i, v := range count {
			if v < m {
				m = v
				mr = uint8(i + 'a')
			}
		}

		for i, v := range count {
			fmt.Printf("%s:%d ", string(rune(i+'a')), v)
		}
		fmt.Println()

		result[i] = mr
	}

	return string(result)
}
