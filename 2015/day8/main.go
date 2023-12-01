package main

import (
	"fmt"
	"strconv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2015/day8/part.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input []string) int {
	total := 0
	for _, v := range input {
		t := unquote(v)
		total += len(v) - len(t)
	}

	return total
}

func unquote(str string) string {
	s, _ := strconv.Unquote(str)

	return s
}

func part2(input []string) interface{} {
	total := 0
	for _, v := range input {
		t := quote(v)
		total += len(t) - len(v)
	}

	return total
}

func quote(str string) string {
	s := strconv.Quote(str)

	return s
}
