package main

import (
	"fmt"
	"go-aoc/pkg/conv"
	"slices"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2025/day5/part1.input")
	//fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

type pair struct {
	start int64
	stop  int64
}

func part1(input []string) any {
	var total int
	var idxs []pair
	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, stop := conv.ToInt64(parts[0]), conv.ToInt64(parts[1])
			idxs = append(idxs, pair{start: start, stop: stop})

			continue
		}

		if isFresh(conv.ToInt64(line), idxs) {
			total++
		}
	}

	return total
}

func isFresh(id int64, idxs []pair) bool {
	for _, idx := range idxs {
		if id >= idx.start && id <= idx.stop {
			return true
		}
	}

	return false
}

func part2(input []string) any {
	var idxs []pair
	for _, line := range input {
		if line == "" {
			break
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, stop := conv.ToInt64(parts[0]), conv.ToInt64(parts[1])
			idxs = append(idxs, pair{start: start, stop: stop})

			continue
		}
	}

	slices.SortFunc(idxs, func(a, b pair) int {
		if a.start <= b.start {
			return -1
		}

		return 1
	})
	fmt.Println(idxs)

	merged := []pair{idxs[0]}
	for i := 1; i < len(idxs); i++ {
		idx := idxs[i]
		last := merged[len(merged)-1]

		if idx.start <= last.stop {
			merged[len(merged)-1] = pair{
				start: last.start,
				stop:  max(idx.stop, last.stop),
			}
			continue
		}

		merged = append(merged, idx)
	}

	fmt.Println(merged)

	var total int64
	for _, idx := range merged {
		total += idx.stop - idx.start + 1
	}
	return total
}
