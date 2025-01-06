package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

var (
	pattern  = "mul\\([0-9]*,[0-9]*\\)"
	pattern2 = "do\\(\\)|don\\'t\\(\\)|mul\\([0-9]*,[0-9]*\\)"
)

func main() {
	input := in.ReadFileAsStringSlice("/2024/day3/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	var sum int
	for _, line := range input {
		r := regexp.MustCompile(pattern)
		matches := r.FindAllString(line, -1)

		for _, m := range matches {
			sum += doMultiply(m)
		}
	}

	return sum
}

func part2(input []string) any {
	var sum int
	enabled := true
	for _, line := range input {
		r := regexp.MustCompile(pattern2)
		matches := r.FindAllString(line, -1)

		for _, m := range matches {
			if !enabled && strings.HasPrefix(m, "do(") {
				enabled = true
				continue
			}

			if enabled && strings.HasPrefix(m, "don't(") {
				enabled = false
				continue
			}

			if enabled && strings.HasPrefix(m, "mul") {
				sum += doMultiply(m)
			}
		}
	}

	return sum
}

func doMultiply(m string) int {
	res1 := strings.Replace(m, "mul(", "", 1)
	res2 := strings.Replace(res1, ")", "", 1)

	parts := strings.Split(res2, ",")

	n1, _ := strconv.Atoi(parts[0])
	n2, _ := strconv.Atoi(parts[1])

	return n1 * n2
}
