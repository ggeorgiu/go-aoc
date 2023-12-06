package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/conv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day2/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) int {
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for _, line := range input {
		parts := strings.Split(line, ": ")
		gameID := conv.ToInt(parts[0][5:]) // getting game ID from "Game x"
		games := strings.Split(parts[1], "; ")

		possible := true
		for _, g := range games {
			if !possible {
				break
			}

			draws := strings.Split(g, ", ")
			for _, d := range draws {
				vals := strings.Split(d, " ")
				c := conv.ToInt(vals[0])
				color := vals[1]

				if c > config[color] {
					possible = false
					break
				}
			}
		}

		if possible {
			sum += gameID
		}
	}

	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		parts := strings.Split(line, ": ")
		games := strings.Split(parts[1], "; ")
		m := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, g := range games {
			draws := strings.Split(g, ", ")
			for _, d := range draws {
				vals := strings.Split(d, " ")
				c := conv.ToInt(vals[0])
				color := vals[1]

				if m[color] < c {
					m[color] = c
				}
			}
		}

		sum += m["red"] * m["blue"] * m["green"]
	}

	return sum
}
