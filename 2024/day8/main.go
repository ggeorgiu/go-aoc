package main

import (
	"fmt"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2024/day8/part1.input")
	fmt.Printf("P1: %v\n", part1(input))

	input = in.ReadFileAsRuneMatrix("/2024/day8/part1.input")
	fmt.Printf("P2: %v\n", part2(input))
}

type point struct {
	x int
	y int
}

func part1(input [][]rune) any {
	count := make(map[point]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				for k := 0; k < len(input); k++ {
					for v := 0; v < len(input[1]); v++ {
						if k == i && v == j || input[k][v] == '#' {
							continue
						}

						if input[k][v] == input[i][j] {
							di := k - i
							dj := v - j

							var ni, nj int
							ni = k + di
							nj = v + dj

							if ni >= 0 && ni < len(input) && nj >= 0 && nj < len(input[0]) {
								count[point{ni, nj}] = true
								if input[ni][nj] == '.' {
									input[ni][nj] = '#'
								}
							}
						}
					}
				}
			}
		}
	}

	return len(count)
}

func part2(input [][]rune) any {
	count := make(map[point]bool)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				for k := 0; k < len(input); k++ {
					for v := 0; v < len(input[1]); v++ {
						if k == i && v == j || input[k][v] == '#' {
							continue
						}

						if input[k][v] == input[i][j] {
							count[point{i, j}] = true
							di := k - i
							dj := v - j

							var ni, nj int
							ni = k + di
							nj = v + dj

							for ni >= 0 && ni < len(input) && nj >= 0 && nj < len(input[0]) {
								count[point{ni, nj}] = true
								if input[ni][nj] == '.' {
									input[ni][nj] = '#'
								}

								ni = ni + di
								nj = nj + dj
							}
						}
					}
				}
			}
		}
	}

	return len(count)
}
