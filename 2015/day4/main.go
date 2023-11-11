package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsString("/2015/day4/part1.input")
	fmt.Printf("P1: %d\n", part1(input))
	fmt.Printf("P2: %d\n", part2(input))
}

func part1(input string) int {
	flag := true
	i := 0
	for flag {
		i++
		res := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
		if strings.HasPrefix(res, "00000") {
			flag = false
		}
	}
	return i
}

func part2(input string) int {
	flag := true
	i := 0
	for flag {
		i++
		res := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
		if strings.HasPrefix(res, "000000") {
			flag = false
		}
	}
	return i
}
