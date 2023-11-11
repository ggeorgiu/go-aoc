package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsString("/2015/day3/part1.input")
	fmt.Printf("P1: %d\n", part1(input))
	fmt.Printf("P2: %d\n", part2(input))
}

type point struct {
	x int
	y int
}

func part1(input string) int {
	var x, y int
	coords := make(map[point]bool)

	coords[point{x, y}] = true
	for _, v := range input {
		switch v {
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		case '^':
			y--
		}

		coords[point{x, y}] = true
	}

	return len(coords)
}

func part2(input string) int {
	ix := []int{0, 0}
	iy := []int{0, 0}
	coords := make(map[point]bool)

	coords[point{ix[0], iy[0]}] = true
	for i, v := range input {
		switch v {
		case '>':
			ix[i%2]++
		case 'v':
			iy[i%2]++
		case '<':
			ix[i%2]--
		case '^':
			iy[i%2]--
		}

		coords[point{ix[i%2], iy[i%2]}] = true
	}

	return len(coords)
}
