package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsIntMatrix("/2024/day2/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input [][]int) any {
	var count int

	for i := 0; i < len(input); i++ {
		var asc bool
		valid := true
		if input[i][0] < input[i][1] {
			asc = true
		}

		for j := 1; j < len(input[i]); j++ {
			var diff int
			if asc {
				diff = input[i][j] - input[i][j-1]
			} else {
				diff = -1 * (input[i][j] - input[i][j-1])
			}

			if diff > 3 || diff <= 0 {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}

	return count
}

func part2(input [][]int) any {
	var count int

	for i := 0; i < len(input); i++ {
		var asc bool
		valid := true
		if input[i][0] < input[i][1] {
			asc = true
		}

		for j := 1; j < len(input[i]); j++ {
			var diff int
			if asc {
				diff = input[i][j] - input[i][j-1]
			} else {
				diff = -1 * (input[i][j] - input[i][j-1])
			}

			if diff > 3 || diff <= 0 {
				valid = false
				break
			}
		}
		if valid {
			count++
			fmt.Println(fmt.Sprintf("%v - %v", "0", input[i]))
			continue
		}

		for w := 0; w < len(input[i]); w++ {
			validd := true
			var val []int

			val = append(val, input[i]...)
			line := append(val[:w], val[w+1:]...)
			var ascc bool
			if line[0] < line[1] {
				ascc = true
			}
			for j := 1; j < len(line); j++ {
				var diff int
				if ascc {
					diff = line[j] - line[j-1]
				} else {
					diff = -1 * (line[j] - line[j-1])
				}

				if diff > 3 || diff <= 0 {
					validd = false
					break
				}
			}

			if validd {
				count++
				break
			}
		}

	}

	return count
}
