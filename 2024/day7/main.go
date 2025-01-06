package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"go-aoc/pkg/gen"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2024/day7/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	config := make(map[int][]string)
	var total int64

	for _, line := range input {
		parts := strings.Split(line, ": ")

		result, _ := strconv.Atoi(parts[0])
		components := strings.Split(parts[1], " ")
		operandLen := len(components) - 1

		if _, ok := config[operandLen]; !ok {
			var res []string
			gen.CartesianProduct("*+", "", operandLen, &res)
			config[operandLen] = res
		}

		val := config[operandLen]
		for _, pos := range val {
			nr1, _ := strconv.Atoi(components[0])
			res := nr1

			for i, c := range pos {
				nr, _ := strconv.Atoi(components[i+1])

				if c == '+' {
					res += nr
				}
				if c == '*' {
					res *= nr
				}
			}
			if res == result {
				total += int64(res)
				break
			}
		}
	}

	return total
}

func part2(input []string) any {
	config := make(map[int][]string)
	var total int64

	for _, line := range input {
		parts := strings.Split(line, ": ")

		result, _ := strconv.Atoi(parts[0])
		components := strings.Split(parts[1], " ")
		operandLen := len(components) - 1

		if _, ok := config[operandLen]; !ok {
			var res []string
			gen.CartesianProduct("*+|", "", operandLen, &res)
			config[operandLen] = res
		}

		val := config[operandLen]
		for _, pos := range val {
			nr1, _ := strconv.Atoi(components[0])
			res := nr1

			for i, c := range pos {
				nr, _ := strconv.Atoi(components[i+1])

				if c == '+' {
					res += nr
				}
				if c == '*' {
					res *= nr
				}
				if c == '|' {
					res = res * int(math.Pow(float64(10), float64(len(strconv.Itoa(nr)))))
					res = res + nr
				}
			}
			if res == result {
				total += int64(res)
				break
			}
		}
	}

	return total
}
