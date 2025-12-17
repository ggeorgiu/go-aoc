package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2025/day7/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input [][]rune) any {
	var starti, startj int

	var found bool
	for i := 0; i < len(input); i++ {
		if found {
			break
		}

		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'S' {
				starti = i
				startj = j
				found = true

				break
			}
		}
	}

	return descend(input, starti+1, startj)
}

func descend(matrix [][]rune, i, j int) int {
	if i == len(matrix) {
		return 0
	}

	if matrix[i][j] == '^' {
		if matrix[i][j-1] == '|' && matrix[i][j+1] == '|' {
			return 0
		}
		if matrix[i][j-1] != '.' {
			return 1 + descend(matrix, i, j+1)
		}
		if matrix[i][j+1] != '.' {
			return 1 + descend(matrix, i, j-1)
		}

		return 1 + descend(matrix, i, j-1) + descend(matrix, i, j+1)
	}

	matrix[i][j] = '|'
	return descend(matrix, i+1, j)
}

func part2(input [][]rune) any {
	var starti, startj int

	var found bool
	for i := 0; i < len(input); i++ {
		if found {
			break
		}

		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'S' {
				starti = i
				startj = j
				found = true

				break
			}
		}
	}

	cache := make([][]int, len(input))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(input[i]))
	}

	var total int
	var lastLine = len(input) - 1
	for j := 0; j < len(input[lastLine]); j++ {
		if input[lastLine][j] == '|' {
			c := ascend(input, lastLine, j, starti, startj, cache)
			total += c

		}
	}

	return total
}

func ascend(matrix [][]rune, i, j int, targeti, targetj int, cache [][]int) int {
	if i < 0 || j < 0 || j > len(matrix[i]) || matrix[i][j] == '.' {
		return 0
	}

	if i == targeti && j == targetj {
		return 1
	}

	if cache[i][j] != 0 {
		return cache[i][j]
	}

	if j > 0 && j < len(matrix[i])-1 {
		if matrix[i][j-1] == '^' && matrix[i][j+1] == '^' {
			v1 := ascend(matrix, i-1, j-1, targeti, targetj, cache)
			v2 := ascend(matrix, i-1, j+1, targeti, targetj, cache)
			v3 := ascend(matrix, i-1, j, targeti, targetj, cache)

			cache[i][j] = v1 + v2 + v3
			return v1 + v2 + v3
		}
	}

	if j > 0 {
		if matrix[i][j-1] == '^' {
			v1 := ascend(matrix, i-1, j-1, targeti, targetj, cache)
			v2 := ascend(matrix, i-1, j, targeti, targetj, cache)
			cache[i][j] = v1 + v2
			return v1 + v2
		}
	}

	if j < len(matrix[i])-1 {
		if matrix[i][j+1] == '^' {
			v1 := ascend(matrix, i-1, j+1, targeti, targetj, cache)
			v2 := ascend(matrix, i-1, j, targeti, targetj, cache)

			cache[i][j] = v1 + v2
			return v1 + v2

		}
	}

	v1 := ascend(matrix, i-1, j, targeti, targetj, cache)
	cache[i][j] = v1

	return v1
}
