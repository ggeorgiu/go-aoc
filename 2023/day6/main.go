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

func part1(input []string) interface{} {
	conf := make(map[string][]int)
	order := []string{"time", "distance"}

	for i, v := range input {
		parts := strings.Split(v, ":")
		for _, v := range strings.Split(parts[1], " ") {
			if v == "" {
				continue
			}

			conf[order[i]] = append(conf[order[i]], conv.ToInt(v))
		}
	}

	sum := 1
	for k, v := range conf["time"] {
		dist := 0
		count := 0
		for i := 0; i <= v; i++ {
			dist = i * (v - i)
			if dist > conf["distance"][k] {
				count++
			}
		}
		sum *= count
	}

	return sum
}

func part2(input []string) interface{} {
	conf := make(map[string]int)
	order := []string{"time", "distance"}

	for i, v := range input {
		parts := strings.Split(v, ":")
		conf[order[i]] = conv.ToInt(strings.ReplaceAll(parts[1], " ", ""))
	}

	dist := 0
	count := 0
	v := conf["time"]
	for i := 0; i <= v; i++ {
		dist = i * (v - i)
		if dist > conf["distance"] {
			count++
		}
	}

	return count
}
