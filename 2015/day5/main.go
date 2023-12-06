package main

import (
	"fmt"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2015/day5/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) int {
	c := 0
	for _, s := range input {
		if containsThreeVowels(s) && hasDuplicateLetter(s) && doesNotContainStrings(s) {
			c++
		}
	}

	return c
}

func part2(input []string) int {
	c := 0
	for _, s := range input {
		if containsTwoPairs(s) && hasDuplicateLetterSeparated(s) {
			c++
		}
	}

	return c
}

func doesNotContainStrings(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}

	for i := range forbidden {
		if strings.Contains(s, forbidden[i]) {
			return false
		}
	}

	return true
}

func hasDuplicateLetter(s string) bool {
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			return true
		}
	}

	return false
}

func containsThreeVowels(s string) bool {
	vowels := "aeiou"
	ct := 0
	for _, c := range strings.Split(s, "") {
		if strings.Contains(vowels, c) {
			ct++
		}
	}

	return ct >= 3
}

func containsTwoPairs(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func hasDuplicateLetterSeparated(s string) bool {
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-2; i++ {
		if ss[i] == ss[i+2] {
			return true
		}
	}

	return false
}
