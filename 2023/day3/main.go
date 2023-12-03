package main

import (
	"fmt"
	"strconv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2023/day3/part1.test")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input [][]rune) interface{} {
	sum := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if !isSign(input[i][j]) {
				continue
			}

			nums := getNeighborNums(input, i, j)
			if len(nums) == 0 {
				continue
			}

			for _, v := range nums {
				sum += v
			}
		}
	}

	return sum
}

func part2(input [][]rune) interface{} {
	var sum int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '*' {
				continue
			}

			nums := getNeighborNums(input, i, j)
			if len(nums) <= 1 {
				continue
			}

			p := 1
			for _, v := range nums {
				p *= v
			}

			sum += p
		}
	}

	return sum
}

type cord struct {
	x int
	y int
}

func getNeighborNums(input [][]rune, x, y int) []int {
	neighs := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	visited := make(map[cord]bool)
	var nrs []int
	for i := 0; i < len(neighs); i++ {
		ni := x + neighs[i][0]
		nj := y + neighs[i][1]

		if !visited[cord{ni, nj}] && isDigit(input[ni][nj]) {
			start, end := nj, nj
			for start > 0 && isDigit(input[ni][start-1]) {
				start--
				visited[cord{ni, start}] = true

			}
			for end < len(input[0])-1 && isDigit(input[ni][end+1]) {
				end++
				visited[cord{ni, end}] = true
			}

			n, _ := strconv.Atoi(string(input[ni][start : end+1]))
			nrs = append(nrs, n)
		}
	}

	return nrs
}

func isSign(r rune) bool {
	return !isDigit(r) && r != '.'
}
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
