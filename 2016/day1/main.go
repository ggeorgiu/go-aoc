package main

import (
	"fmt"
	"math"
	"strconv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsCSVLine("/2016/day1/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

var vals = []int{1, 1, -1, -1}

func part1(input []string) interface{} {
	var pos = []int{0, 0}

	orientation := 0
	for _, v := range input {
		switch v[0] {
		case 'L':
			orientation = (orientation - 1) % 4
		case 'R':
			orientation = (orientation + 1) % 4
		}
		if orientation == -1 {
			orientation = 3
		}

		snr := v[1:]
		nr, _ := strconv.Atoi(snr)
		pos[orientation%2] += nr * vals[orientation]
	}

	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

type point struct {
	a int
	b int
}

func part2(input []string) interface{} {
	orientation := 0
	var pos = []int{0, 0}

	mem := map[point]bool{
		{0, 0}: true,
	}

	for _, v := range input {
		switch v[0] {
		case 'L':
			orientation = (orientation - 1) % 4
		case 'R':
			orientation = (orientation + 1) % 4
		}
		if orientation == -1 {
			orientation = 3
		}

		snr := v[1:]
		nr, _ := strconv.Atoi(snr)

		for i := 1; i <= nr; i++ {
			pos[orientation%2] += 1 * vals[orientation]
			if mem[point{pos[0], pos[1]}] {
				return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
			}
			mem[point{pos[0], pos[1]}] = true
		}

	}

	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}
