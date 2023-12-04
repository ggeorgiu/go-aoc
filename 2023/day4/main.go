package main

import (
	"fmt"
	"math"
	"strings"

	"go-aoc/pkg/conv"
	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day4/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input []string) interface{} {
	var sum int
	arr := make([]int, len(input))

	for i, v := range input {
		count := 0
		v = v[10:]
		parts := strings.Split(v, " | ")

		winning := strings.Split(parts[0], " ")
		w := conv.ToIntSlice(winning)

		got := strings.Split(parts[1], " ")
		g := conv.ToIntSlice(got)

		for iw := 0; iw < len(w); iw++ {
			for jg := 0; jg < len(g); jg++ {
				if g[jg] == w[iw] {
					count++
				}
			}
		}

		arr[i] = count
	}

	for _, v := range arr {
		if v != 0 {
			sum += int(math.Pow(2, float64(v-1)))
		}
	}

	return sum
}

func part2(input []string) interface{} {
	arr := make([]int, len(input))
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}

	sum := 0
	for i, v := range input {
		v = v[10:]
		parts := strings.Split(v, " | ")

		winning := strings.Split(parts[0], " ")
		w := conv.ToIntSlice(winning)

		got := strings.Split(parts[1], " ")
		g := conv.ToIntSlice(got)

		count := 0
		for iw := 0; iw < len(w); iw++ {
			for jg := 0; jg < len(g); jg++ {
				if g[jg] == w[iw] {
					count++
				}
			}
		}

		k := 1
		for k <= count {
			arr[i+k] += arr[i]
			k++
		}
	}

	for _, v := range arr {
		sum += v
	}

	return sum
}
