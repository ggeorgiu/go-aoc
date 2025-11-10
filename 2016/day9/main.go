package main

import (
	"fmt"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsString("/2016/day9/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(line string) any {
	var result strings.Builder
	for i := 0; i < len(line); i++ {
		if line[i] == '(' {
			var group []uint8
			var k = i + 1
			for line[k] != ')' {
				group = append(group, line[k])
				k++
			}

			sgroup := string(group)
			parts := strings.Split(sgroup, "x")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			part := line[k+1 : k+1+x]

			for y > 0 {
				result.WriteString(part)
				y--
			}
			i = k + x
		} else {
			result.WriteRune(rune(line[i]))
		}
	}

	return result.Len()
}

func part2(line string) any {
	return compute(line)
}

func compute(part string) int {
	if !hasGroup(part) {
		return len(part)
	}

	var pref int
	for i := 0; i < len(part); i++ {
		if part[i] == '(' {
			break
		}
		pref++
	}

	var x, y int
	var group []uint8
	for i := pref + 1; part[i] != ')'; i++ {
		group = append(group, part[i])
	}

	sgroup := string(group)
	parts := strings.Split(sgroup, "x")
	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])

	from := pref + len(group) + 2
	to := from + x

	return pref + y*compute(part[from:to]) + compute(part[to:])
}

func hasGroup(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			return true
		}
	}

	return false
}
