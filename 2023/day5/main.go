package main

import (
	"fmt"
	"math"
	"strings"

	"go-aoc/pkg/conv"

	"go-aoc/pkg/in"
)

type interval struct {
	start      int
	end        int
	startValue int
}

// Brute force
func main() {
	input := in.ReadFileAsStringSlice("/2023/day5/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) interface{} {
	seeds, conf := parseSeedsAndConfig(input)

	workingMap := make(map[int]int)
	for _, v := range seeds {
		workingMap[v] = v
	}

	order := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	for _, o := range order {
		for k, v := range workingMap {
			for _, i := range conf[o] {
				if v >= i.start && v <= i.end {
					workingMap[k] = i.startValue + v - i.start
				}
			}
		}
	}

	mx := math.MaxInt
	for _, v := range workingMap {
		if v < mx {
			mx = v
		}
	}

	return mx
}

func part2(input []string) interface{} {
	seeds, conf := parseSeedsAndConfig(input)

	mx := math.MaxInt
	order := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	var current int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			current = j
			for _, o := range order {
				for _, iv := range conf[o] {
					if current >= iv.start && current <= iv.end {
						current = iv.startValue + current - iv.start
						break
					}
				}
			}

			if current < mx {
				mx = current
			}
		}
	}

	return mx
}

func parseSeedsAndConfig(input []string) ([]int, map[string][]interval) {
	m := make(map[string][]interval)
	var seeds []int

	for _, line := range input {
		current := ""

		switch {
		case strings.Contains(line, "seeds:"):
			parts := strings.Split(line, ": ")
			for _, v := range strings.Split(parts[1], " ") {
				is := conv.ToInt(v)
				seeds = append(seeds, is)
			}

			continue
		case strings.Contains(line, "map:"):
			parts := strings.Split(line, " ")
			ft := strings.Split(parts[0], "-to-")

			current = ft[1]
			continue

		case line == "":
			continue
		}

		parts := strings.Split(line, " ")
		destCorresp := conv.ToInt(parts[0])
		sourceStart := conv.ToInt(parts[1])
		size := conv.ToInt(parts[2])

		m[current] = append(m[current], interval{sourceStart, sourceStart + size - 1, destCorresp})
	}

	return seeds, m
}
