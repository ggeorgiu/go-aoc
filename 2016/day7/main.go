package main

import (
	"fmt"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day7/part1.input")
	//fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part2(input []string) any {
	var count int

	for _, line := range input {
		var inside bool

		var abasI []string
		var abasO []string

		for i := 0; i < len(line)-2; i++ {
			if line[i] == '[' {
				inside = true
				continue
			}
			if line[i] == ']' {
				inside = false
				continue
			}

			if line[i] == line[i+2] && line[i] != line[i+1] && line[i+1] != '[' && line[i+1] != ']' {
				if inside {
					abasI = append(abasI, line[i:i+3])
					continue
				}
				abasO = append(abasO, line[i:i+3])
			}
		}

		for _, ki := range abasI {
			var counted bool
			for _, ko := range abasO {
				if ki[0] == ko[1] && ki[1] == ko[0] {
					count++
					counted = true
					break
				}
			}
			if counted {
				break
			}
		}
	}

	return count
}

func part1(input []string) any {
	var count int

	for _, line := range input {
		var inside bool
		var hadInside bool
		var found bool

		for i := 0; i < len(line)-3; i++ {
			if line[i] == '[' {
				inside = true
				continue
			}
			if line[i] == ']' {
				inside = false
				continue
			}

			if line[i] == line[i+3] && line[i+1] == line[i+2] && line[i] != line[i+1] {
				found = true
				if inside {
					hadInside = true
					break
				}
			}
		}

		if found && !hadInside {
			//fmt.Println(line)
			count++
		}
	}

	return count
}
