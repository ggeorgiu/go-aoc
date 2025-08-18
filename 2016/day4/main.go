package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day4/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) any {
	var sum int
	for _, line := range input {
		dict := make(map[rune]int)

		parts := strings.Split(line, "-")
		for x, part := range parts {
			if x == len(parts)-1 {
				break
			}

			for _, c := range part {
				dict[c] = dict[c] + 1
			}
		}

		keys := make([]rune, 0, len(dict))
		for key := range dict {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			if dict[keys[i]] == dict[keys[j]] {
				return keys[i] < keys[j]
			}

			return dict[keys[i]] > dict[keys[j]]
		})

		id, _ := strconv.Atoi(parts[len(parts)-1][:3])
		cs := parts[len(parts)-1][4:9]
		ks := keys[:5]

		if cs == string(ks) {
			sum += id
		}
	}

	return sum
}

var abc = "abcdefghijklmnopqrstuvwxyx"

func part2(input []string) any {
	for _, line := range input {
		dict := make(map[rune]int)

		parts := strings.Split(line, "-")
		for x, part := range parts {
			if x == len(parts)-1 {
				break
			}

			for _, c := range part {
				dict[c] = dict[c] + 1
			}
		}

		keys := make([]rune, 0, len(dict))
		for key := range dict {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			if dict[keys[i]] == dict[keys[j]] {
				return keys[i] < keys[j]
			}

			return dict[keys[i]] > dict[keys[j]]
		})

		id, _ := strconv.Atoi(parts[len(parts)-1][:3])
		cs := parts[len(parts)-1][4:9]
		ks := keys[:5]

		if cs != string(ks) {
			continue
		}

		out := make([]rune, 0, len(line))
		for _, r := range line {
			if r >= '0' && r <= '9' {
				break
			}
			if r == '-' {
				out = append(out, ' ')
				continue
			}

			res := (int(r-'a') + (id % 26)) % 26
			out = append(out, rune(abc[res]))
		}

		if strings.Contains(string(out), "northpole") {
			return id
		}
	}

	return 0
}
