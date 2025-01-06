package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day2/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

var dial = [][]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

func part1(input []string) any {
	posL := 1
	posC := 1

	var sb strings.Builder
	for _, line := range input {
		for _, c := range line {
			nextL := posL
			nextC := posC

			switch c {
			case 'U':
				nextL--
			case 'D':
				nextL++
			case 'L':
				nextC--
			case 'R':
				nextC++
			}

			if nextL >= 0 && nextL <= 2 {
				posL = nextL
			}
			if nextC >= 0 && nextC <= 2 {
				posC = nextC
			}
		}

		sb.WriteString(dial[posL][posC])
	}

	return sb.String()
}

var dial2 = [][]string{
	{"_", "_", "_", "_", "_", "_", "_"},
	{"_", "_", "_", "1", "_", "_", "_"},
	{"_", "_", "2", "3", "4", "_", "_"},
	{"_", "5", "6", "7", "8", "9", "_"},
	{"_", "_", "A", "B", "C", "_", "_"},
	{"_", "_", "_", "D", "_", "_", "_"},
	{"_", "_", "_", "_", "_", "_", "_"},
}

func part2(input []string) any {
	posL := 3
	posC := 3

	var sb strings.Builder
	for _, line := range input {
		for _, c := range line {
			nextL := posL
			nextC := posC

			switch c {
			case 'U':
				nextL--
			case 'D':
				nextL++
			case 'L':
				nextC--
			case 'R':
				nextC++
			}

			if dial2[nextL][nextC] != "_" {
				posL = nextL
				posC = nextC
			}
		}
		sb.WriteString(dial2[posL][posC])
	}

	return sb.String()
}
