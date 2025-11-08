package main

import (
	"fmt"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

var screen [][]bool

const (
	lines = 6
	cols  = 50
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day8/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
}

func part1(input []string) any {
	screen = make([][]bool, lines)
	for i := range screen {
		screen[i] = make([]bool, cols)
	}

	for _, instr := range input {
		switch {
		case strings.HasPrefix(instr, "rect"):
			parts := strings.Split(instr[5:], "x")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			drawRect(x, y)
		default:
			parts := strings.Split(instr, " ")

			x, _ := strconv.Atoi(parts[2][2:])
			y, _ := strconv.Atoi(parts[4])
			switch {
			case parts[1] == "row":
				shiftLine(x, y)
			default:
				shiftColumn(x, y)
			}

		}
	}
	printScreen()

	return count()
}

func drawRect(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			screen[i][j] = true
		}
	}
}

func shiftLine(idx int, by int) {
	for k := 0; k < by; k++ {

		val := screen[idx][len(screen[idx])-1]
		for i := len(screen[idx]) - 1; i > 0; i-- {
			screen[idx][i] = screen[idx][i-1]
		}

		screen[idx][0] = val
	}
}

func shiftColumn(idx int, by int) {
	for k := 0; k < by; k++ {

		val := screen[len(screen)-1][idx]
		for i := len(screen) - 1; i > 0; i-- {
			screen[i][idx] = screen[i-1][idx]
		}

		screen[0][idx] = val
	}
}

func printScreen() {
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func count() int {
	var total int
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				total++
			}
		}
	}
	return total
}
