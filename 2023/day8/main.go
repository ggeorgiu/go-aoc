package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/algs"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day8/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

type node struct {
	left  string
	right string
}

func part1(input []string) any {
	directions := input[0]

	config := make(map[string]node)
	for i := 2; i < len(input); i++ {
		parts := strings.Split(input[i], " = ")
		key := parts[0]

		lr := strings.Split(parts[1], ", ")
		left := strings.TrimPrefix(lr[0], "(")
		right := strings.TrimSuffix(lr[1], ")")

		config[key] = node{
			left:  left,
			right: right,
		}
	}

	current := "AAA"
	di := 0
	step := 0
	for current != "ZZZ" {
		n := config[current]
		switch directions[di] {
		case 'R':
			current = n.right
		case 'L':
			current = n.left
		}
		step++
		di++
		if di == len(directions) {
			di = 0
		}
	}

	return step
}

func part2(input []string) any {
	directions := input[0]
	config := make(map[string]node)
	var current []string
	for i := 2; i < len(input); i++ {
		parts := strings.Split(input[i], " = ")
		key := parts[0]
		lr := strings.Split(parts[1], ", ")

		left := strings.TrimPrefix(lr[0], "(")
		right := strings.TrimSuffix(lr[1], ")")

		config[key] = node{
			left:  left,
			right: right,
		}

		if key[2] == 'A' {
			current = append(current, key)
		}
	}

	di := 0
	steps := make([]int, len(current))
	cycles := make([]int, len(current))

	for !allPopulated(cycles) {
		for i := 0; i < len(current); i++ {
			n := config[current[i]]
			switch {
			case directions[di] == 'R':
				current[i] = n.right
				current[i] = n.left
			case directions[di] == 'L':
				current[i] = n.left
			case cycles[i] == 0 && current[i][2] != 'Z':
				steps[i]++
			case cycles[i] == 0 && current[i][2] == 'Z':
				cycles[i] = steps[i] + 1
			}
		}

		di++
		if di == len(directions) {
			di = 0
		}
	}

	return algs.LCM(cycles...)
}

func allPopulated(nodes []int) bool {
	for _, v := range nodes {
		if v == 0 {
			return false
		}
	}

	return true
}
