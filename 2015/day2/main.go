package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/conv"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2015/day2/part1.input")
	fmt.Printf("P1: %d\n", part1(input))
	fmt.Printf("P2: %d\n", part2(input))
}

func part1(input []string) int {
	var total int
	for _, line := range input {
		parts := strings.Split(line, "x")
		l, w, h := conv.ToInt(parts[0]), conv.ToInt(parts[1]), conv.ToInt(parts[2])

		total += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}

	return total
}

func part2(input []string) int {
	var total int
	for _, line := range input {
		parts := strings.Split(line, "x")
		l, w, h := conv.ToInt(parts[0]), conv.ToInt(parts[1]), conv.ToInt(parts[2])

		total += l*w*h + 2*(l+w+h) - 2*max(l, w, h)
	}

	return total
}
