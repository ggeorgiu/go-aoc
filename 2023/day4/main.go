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
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	var sum int
	arr := make([]int, len(input))

	p := strings.Index(input[0], ":") + 2
	for i, line := range input {
		w, g := parseCards(line, p)

		count := 0
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

func part2(input []string) any {
	cardCount := make([]int, len(input))
	for i := 0; i < len(cardCount); i++ {
		cardCount[i] = 1
	}

	sum := 0
	p := strings.Index(input[0], ":") + 2
	for i, line := range input {
		w, g := parseCards(line, p)

		count := 0
		for iw := 0; iw < len(w); iw++ {
			for jg := 0; jg < len(g); jg++ {
				if g[jg] == w[iw] {
					count++
				}
			}
		}

		cIdx := 1
		for cIdx <= count {
			cardCount[i+cIdx] += cardCount[i]
			cIdx++
		}
	}

	for _, v := range cardCount {
		sum += v
	}

	return sum
}

func parseCards(line string, p int) ([]int, []int) {
	line = line[p:] // cut the "card x:" prefix
	parts := strings.Split(line, " | ")

	winning := strings.Split(parts[0], " ")
	w := conv.ToIntSlice(winning)

	got := strings.Split(parts[1], " ")
	g := conv.ToIntSlice(got)

	return w, g
}
