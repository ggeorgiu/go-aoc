package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day1/part.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input []string) int {
	var sum int
	for _, v := range input {
		var n int
		for i := 0; i < len(v); i++ {
			if v[i] >= '0' && v[i] <= '9' {
				n = (n * 10) + (int(v[i]) - '0')
				break
			}
		}

		for i := len(v) - 1; i >= 0; i-- {
			if v[i] >= '0' && v[i] <= '9' {
				n = (n * 10) + (int(v[i]) - '0')
				break
			}
		}

		sum += n
	}

	return sum
}

func part2(input []string) int {
	var digits = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	digToInt := map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	var sum int
	for _, v := range input {
		f := make(map[int]string)
		l := make(map[int]string)
		var n int

		for _, d := range digits {
			fi := strings.Index(v, d)
			if fi != -1 {
				f[fi] = d
			}

			li := strings.LastIndex(v, d)
			if li != -1 {
				l[li] = d
			}
		}

		for i := 0; i < len(v); i++ {
			if val, ok := f[i]; ok {
				n = n*10 + digToInt[val]
				break
			}
		}

		for i := len(v) - 1; i >= 0; i-- {
			if val, ok := l[i]; ok {
				n = n*10 + digToInt[val]
				break
			}
		}

		sum += n
	}

	return sum
}
