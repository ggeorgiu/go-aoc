package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsString("/2015/day1/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input string) int {
	var total int
	for _, r := range input {
		switch r {
		case '(':
			total++
		case ')':
			total--
		}
	}

	return total
}

func part2(input string) any {
	var c int
	for i, r := range input {
		switch r {
		case '(':
			c++
		case ')':
			c--
		}

		if c == -1 {
			return i + 1
		}
	}

	return 0
}
