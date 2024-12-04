package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2024/day1/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) interface{} {
	var arr1 []int
	var arr2 []int

	for _, line := range input {
		parts := strings.Fields(line)

		nr1, _ := strconv.Atoi(parts[0])
		nr2, _ := strconv.Atoi(parts[1])

		arr1 = append(arr1, nr1)
		arr2 = append(arr2, nr2)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	var diff int
	for i := 0; i < len(arr1); i++ {
		d := arr2[i] - arr1[i]
		if d < 0 {
			d = d * -1
		}

		diff += d
	}

	return diff
}

func part2(input []string) interface{} {
	var stat = make(map[int]int, 100_000)
	var arr1 []int

	for _, line := range input {
		parts := strings.Fields(line)

		nr1, _ := strconv.Atoi(parts[0])
		nr2, _ := strconv.Atoi(parts[1])

		arr1 = append(arr1, nr1)
		stat[nr2] = stat[nr2] + 1
	}

	var score int
	for _, v := range arr1 {
		score += v * stat[v]
	}

	return score
}
