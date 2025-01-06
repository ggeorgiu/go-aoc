package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2024/day4/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input [][]rune) any {
	var count int

	count += countByLine(input)
	count += countByColumn(input)
	count += countByMainDiagonal(input)
	count += countBySecondDiagonal(input)

	return count
}

func countByMainDiagonal(input [][]rune) int {
	var total int

	for i := 0; i < len(input); i++ {
		var val strings.Builder
		row := i
		for j := 0; row >= 0 && j < len(input[0]); j++ {
			val.WriteRune(input[row][j])
			row--
		}
		total += countXmas(val.String())
	}

	for j := 1; j < len(input[0]); j++ {
		var val strings.Builder
		col := j
		for i := len(input) - 1; i > 0 && col < len(input[0]); i-- {
			val.WriteRune(input[i][col])
			col++
		}
		total += countXmas(val.String())
	}

	return total
}

func countBySecondDiagonal(input [][]rune) int {
	var total int

	for i := 0; i < len(input); i++ {
		var val strings.Builder
		row := i
		for j := len(input[0]) - 1; row >= 0 && j >= 0; j-- {
			val.WriteRune(input[row][j])
			row--
		}
		total += countXmas(val.String())
	}

	for j := len(input[0]) - 2; j >= 0; j-- {
		var val strings.Builder
		col := j
		for i := len(input) - 1; i > 0 && col >= 0; i-- {
			val.WriteRune(input[i][col])
			col--
		}
		total += countXmas(val.String())
	}

	return total
}

func countXmas(val string) int {
	count1 := strings.Count(val, "XMAS")
	count2 := strings.Count(val, "SAMX")

	return count1 + count2
}

func countByColumn(input [][]rune) int {
	var total int
	for i := 0; i < len(input); i++ {
		var val strings.Builder
		for j := 0; j < len(input[i]); j++ {
			val.WriteRune(input[j][i])
		}

		total += countXmas(val.String())
	}

	return total
}

func countByLine(input [][]rune) int {
	var total int
	for i := 0; i < len(input); i++ {
		var val strings.Builder
		for j := 0; j < len(input[i]); j++ {
			val.WriteRune(input[i][j])
		}

		check := val.String()

		count1 := strings.Count(check, "XMAS")
		count2 := strings.Count(check, "SAMX")

		total = total + count1 + count2
	}

	return total
}

func part2(input [][]rune) any {
	var total int
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			if input[i][j] == 'A' {
				if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' {
					total++
				}

				if input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' && input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' {
					total++
				}

				if input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S' && input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M' {
					total++
				}

				if input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M' {
					total++
				}
			}
		}
	}

	return total
}
