package main

import (
	"fmt"
	"math"
	"strings"

	"go-aoc/pkg/conv"

	"go-aoc/pkg/in"
)

type interval struct {
	start int
	end   int
	sv    int
}

func main() {
	input := in.ReadFileAsStringSlice("/2023/day5/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) interface{} {
	m := make(map[string][]interval)
	sm := make(map[int]int)

	var seeds []int
	to := ""
	for _, v := range input {
		if strings.Contains(v, "seeds:") {
			parts := strings.Split(v, ": ")
			for _, v := range strings.Split(parts[1], " ") {
				is := conv.ToInt(v)
				seeds = append(seeds, is)
				sm[is] = is
			}

			continue
		}
		if strings.Contains(v, "map:") {
			parts := strings.Split(v, " ")
			ft := strings.Split(parts[0], "-to-")

			to = ft[1]
			continue
		}

		if v == "" {
			continue
		}

		parts := strings.Split(v, " ")
		destCoresp := conv.ToInt(parts[0])
		sourceStart := conv.ToInt(parts[1])
		size := conv.ToInt(parts[2])

		m[to] = append(m[to], interval{sourceStart, sourceStart + size - 1, destCoresp})
	}

	fmt.Println(m)

	order := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	for _, o := range order {
		for k, v := range sm {
			for _, i := range m[o] {
				if v >= i.start && v <= i.end {
					sm[k] = i.sv + v - i.start
				}
			}
		}
	}

	mx := math.MaxInt
	for _, v := range sm {
		if v < mx {
			mx = v
		}
	}

	return mx
}

func part2(input []string) interface{} {
	m := make(map[string][]interval)

	var seeds []int
	to := ""
	for _, v := range input {
		if strings.Contains(v, "seeds:") {
			parts := strings.Split(v, ": ")
			for _, v := range strings.Split(parts[1], " ") {
				is := conv.ToInt(v)
				seeds = append(seeds, is)
			}

			continue
		}
		if strings.Contains(v, "map:") {
			parts := strings.Split(v, " ")
			ft := strings.Split(parts[0], "-to-")

			to = ft[1]
			continue
		}

		if v == "" {
			continue
		}

		parts := strings.Split(v, " ")
		destCoresp := conv.ToInt(parts[0])
		sourceStart := conv.ToInt(parts[1])
		size := conv.ToInt(parts[2])

		m[to] = append(m[to], interval{sourceStart, sourceStart + size - 1, destCoresp})
	}

	mx := math.MaxInt
	current := 0
	order := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			current = j
			for _, o := range order {
				for _, iv := range m[o] {
					if current >= iv.start && current <= iv.end {
						current = iv.sv + current - iv.start
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
