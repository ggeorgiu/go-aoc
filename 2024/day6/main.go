package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsRuneMatrix("/2024/day6/part1.input")
	fmt.Printf("P1: %v\n", part1(input))

	input = in.ReadFileAsRuneMatrix("/2024/day6/part1.input")
	fmt.Printf("P2: %v\n", part2(input))
}

var (
	dirX = []int{-1, 0, 1, 0}
	dirY = []int{0, 1, 0, -1}
)

func part1(input [][]rune) interface{} {
	var posX, posY int

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '^' {
				posX = i
				posY = j
				break
			}
		}
		if posX != 0 {
			break
		}
	}
	input[posX][posY] = 'X'

	count := 1
	dir := 0
	for posX != 0 && posX != len(input)-1 && posY != 0 && posY != len(input[0])-1 {
		if input[posX+dirX[dir]][posY+dirY[dir]] == '#' {
			dir++
			if dir == len(dirX) {
				dir = 0
			}
		}

		posX += dirX[dir]
		posY += dirY[dir]
		if input[posX][posY] != 'X' {
			count++
			input[posX][posY] = 'X'
		}
	}

	return count
}

type point struct {
	i   int
	j   int
	dir int
}

func part2(input [][]rune) interface{} {
	var possX, possY int

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '^' {
				possX = i
				possY = j
				break
			}
		}
		if possX != 0 {
			break
		}
	}

	input[possX][possY] = '.'

	count := 0
	var visited map[point]bool
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' || i == possX && j == possY {
				continue
			}
			visited = map[point]bool{}
			dir := 0

			posX := possX
			posY := possY

			input[i][j] = '#'
			for posX+dirX[dir] >= 0 && posX+dirX[dir] < len(input) && posY+dirY[dir] >= 0 && posY+dirY[dir] < len(input[0]) {
				if input[posX+dirX[dir]][posY+dirY[dir]] == '#' {
					dir++
					if dir == len(dirX) {
						dir = 0
					}
					continue
				}

				if _, ok := visited[point{posX, posY, dir}]; ok {
					count++
					break
				}
				visited[point{posX, posY, dir}] = true

				posX = posX + dirX[dir]
				posY = posY + dirY[dir]

			}

			input[i][j] = '.'
		}
	}

	return count
}
