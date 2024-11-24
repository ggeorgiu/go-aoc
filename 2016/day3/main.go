package main

import (
	"fmt"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day3/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) interface{} {
	var possible int

	for _, line := range input {
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		if a+b <= c || a+c <= b || b+c <= a {
			continue
		}

		possible++
	}

	return possible
}

func part2(input []string) interface{} {
	var possible int
	var triangles = make([][]int, 3)
	for i := 0; i < 3; i++ {
		triangles[i] = make([]int, 3)
	}

	var count int
	for _, line := range input {
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])
		triangles[count][0] = a
		triangles[count][1] = b
		triangles[count][2] = c

		if count != 2 {
			count++
			continue
		}

		count = 0
		for i := 0; i < 3; i++ {
			a0 := triangles[0][i]
			a1 := triangles[1][i]
			a2 := triangles[2][i]
			if a0+a1 <= a2 || a0+a2 <= a1 || a1+a2 <= a0 {
				continue
			}
			possible++
		}
	}

	return possible
}
