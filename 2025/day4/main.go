package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2025/day4/part1.input")
	//fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input [][]rune) any {
	var dirx = []int{-1, -1, -1, 0, 1, 1, 1, 0}
	var diry = []int{-1, 0, 1, 1, 1, 0, -1, -1}

	inBounds := func(a, b int) bool {
		return a >= 0 && a < len(input) && b >= 0 && b < len(input[a])
	}

	var total int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != '@' {
				continue
			}

			var count int

			for d := 0; d < 8; d++ {
				x, y := i+dirx[d], j+diry[d]
				if !inBounds(x, y) {
					continue
				}
				if input[x][y] == '@' {
					count++
				}
			}

			if count < 4 {
				total++
			}
		}
	}

	return total
}

func part2(input [][]rune) any {
	var dirx = []int{-1, -1, -1, 0, 1, 1, 1, 0}
	var diry = []int{-1, 0, 1, 1, 1, 0, -1, -1}

	inBounds := func(a, b int) bool {
		return a >= 0 && a < len(input) && b >= 0 && b < len(input[a])
	}

	removed := true

	workingOn := input
	newConfig := input

	var total int
	for removed {
		removed = false
		workingOn = newConfig
		for i := 0; i < len(workingOn); i++ {
			for j := 0; j < len(workingOn[i]); j++ {
				if workingOn[i][j] != '@' {
					continue
				}

				var count int

				for d := 0; d < 8; d++ {
					x, y := i+dirx[d], j+diry[d]
					if !inBounds(x, y) {
						continue
					}
					if workingOn[x][y] == '@' {
						count++
					}
				}

				if count < 4 {
					newConfig[i][j] = '.'
					total++
					removed = true
				}
			}
		}
	}
	return total
}
