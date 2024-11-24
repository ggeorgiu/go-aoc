package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

var config = map[rune][][]rune{
	'|': {{}, {'7', 'F', '|', 'S'}, {}, {'J', 'L', '|', 'S'}},
	'-': {{'L', 'F', '-', 'S'}, {}, {'J', '7', '-', 'S'}, {}},
	'L': {{}, {'7', 'F', '|', 'S'}, {'7', '-', 'S'}, {}},
	'J': {{'L', 'F', '-', 'S'}, {'7', 'F', '|', 'S'}, {}, {}},
	'7': {{'L', 'F', '-', 'S'}, {}, {}, {'J', 'L', '|', 'S'}},
	'F': {{}, {}, {'7', '-', 'S'}, {'J', 'L', '|', 'S'}},
}

func main() {
	input := in.ReadFileAsRuneMatrix("/2023/day10/part1.test")
	fmt.Printf("P1: %v\n", part1(input))
}

func part1(input [][]rune) interface{} {
	var si, sj int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == 'S' {
				si = i
				sj = j
			}
		}
	}

	neighs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	count := 0
	for {
		dir := 0
		for ni := 0; ni < len(neighs); ni++ {
			ci := si + neighs[ni][0]
			cj := sj + neighs[ni][1]
			if !inBounds(ci, cj, len(neighs)) {
				dir++
				continue
			}

			if v, ok := config[input[ci][cj]]; ok {
				for k := 0; k < len(v[dir]); k++ {
					if v[dir][k] == input[ci][cj] {
						if input[si][sj] != 'S' {
							input[si][sj] = '.'
						}
						count++
						break
					}
				}
			}

			if input[si][sj] == '.' || input[si][sj] == 'S' {
				si = ci
				sj = cj
				break
			}
			dir++
		}

		if input[si][sj] == 'S' {
			break
		}
	}

	return count
}

func inBounds(ci int, cj int, n int) bool {
	return ci > -1 && cj > -1 && ci < n && cj < n
}

func part2(input []string) interface{} {
	return ""
}
