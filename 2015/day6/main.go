package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

const gridSize = 1000

func main() {
	input := in.ReadFileAsStringSlice("/2015/day6/part1.input")
	fmt.Printf("P1: %d\n", part1(input))
	fmt.Printf("P2: %d\n", part2(input))
}

type point struct {
	x int
	y int
}

func part1(input []string) int {
	m := map[point]int{}

	for _, s := range input {
		if strings.HasPrefix(s, "turn on") {
			p1, p2 := getCoords(s, "turn on ")
			turnOn(m, p1, p2)
		} else if strings.HasPrefix(s, "turn off") {
			p1, p2 := getCoords(s, "turn off ")
			turnOff(m, p1, p2)
		} else if strings.HasPrefix(s, "toggle") {
			p1, p2 := getCoords(s, "toggle ")
			toggle(m, p1, p2)
		}
	}

	return count(m)
}

func turnOn(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = 1
		}
	}
}

func turnOff(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = 0
		}
	}
}

func toggle(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] = int(math.Abs(float64(m[point{i, j}] - 1)))
		}
	}
}

func part2(input []string) int {
	m := map[point]int{}

	for _, s := range input {
		if strings.HasPrefix(s, "turn on") {
			p1, p2 := getCoords(s, "turn on ")
			turnOnP2(m, p1, p2)
		} else if strings.HasPrefix(s, "turn off") {
			p1, p2 := getCoords(s, "turn off ")
			turnOffP2(m, p1, p2)
		} else if strings.HasPrefix(s, "toggle") {
			p1, p2 := getCoords(s, "toggle ")
			toggleP2(m, p1, p2)
		}
	}

	return count(m)
}

func turnOnP2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] += 1
		}
	}
}

func turnOffP2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			if v, _ := m[point{i, j}]; v > 0 {
				m[point{i, j}] -= 1
			}
		}
	}
}

func toggleP2(m map[point]int, point1, point2 point) {
	for i := point1.x; i <= point2.x; i++ {
		for j := point1.y; j <= point2.y; j++ {
			m[point{i, j}] += 2
		}
	}
}

func getCoords(s string, pref string) (point, point) {
	s = strings.ReplaceAll(s, pref, "")
	split := strings.Split(s, " through ")
	xs := strings.Split(split[0], ",")
	ys := strings.Split(split[1], ",")

	x1, _ := strconv.Atoi(xs[0])
	y1, _ := strconv.Atoi(xs[1])
	x2, _ := strconv.Atoi(ys[0])
	y2, _ := strconv.Atoi(ys[1])

	return point{x1, y1}, point{x2, y2}
}

func count(m map[point]int) int {
	c := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			v, _ := m[point{i, j}]
			c += v
		}
	}

	return c
}
