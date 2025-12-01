package main

import (
	"fmt"
	"go-aoc/pkg/conv"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2025/day1/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	var count int
	at := 50
	for _, line := range input {
		dir := line[0]
		quan := conv.ToInt(line[1:])

		switch dir {
		case 'L':
			at = at - quan
			if at < 0 {
				at = (100 + at) % 100
			}
		case 'R':
			at = at + quan
			if at > 99 {
				at = at % 100
			}
		}

		if at == 0 {
			count++
		}

	}
	return count
}

func part2(input []string) any {
	var count int
	at := 50
	for _, line := range input {
		dir := line[0]
		quan := conv.ToInt(line[1:])

		switch dir {
		case 'L':
			for quan > 0 {
				at--
				if at == 0 {
					count++
				}
				if at < 0 {
					at = 99
				}
				quan--
			}
		case 'R':
			for quan > 0 {
				at++
				if at == 100 {
					at = 0
				}
				if at == 0 {
					count++
				}
				quan--
			}
		}
	}

	return count
}
