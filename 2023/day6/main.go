package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/conv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day6/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	conf := make(map[string][]int)
	order := []string{"time", "distance"}

	for i, line := range input {
		parts := strings.Split(line, ":")
		for _, v := range strings.Split(parts[1], " ") {
			if v == "" {
				continue
			}

			conf[order[i]] = append(conf[order[i]], conv.ToInt(v))
		}
	}

	sum := 1
	for k, time := range conf["time"] {
		dist := 0
		count := 0
		for i := 0; i <= time; i++ {
			dist = i * (time - i)
			if dist > conf["distance"][k] {
				count++
			}
		}
		sum *= count
	}

	return sum
}

func part2(input []string) any {
	conf := make(map[string]int)
	order := []string{"time", "distance"}

	for i, line := range input {
		parts := strings.Split(line, ":")
		conf[order[i]] = conv.ToInt(strings.ReplaceAll(parts[1], " ", ""))
	}

	dist := 0
	count := 0
	time := conf["time"]
	for i := 0; i <= time; i++ {
		dist = i * (time - i)
		if dist > conf["distance"] {
			count++
		}
	}

	return count
}
