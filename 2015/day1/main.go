package main

import (
	"fmt"
	"os"
)

func main() {
	input := ReadFileAsString("/day1/part1.input")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	var total int
	for _, r := range input {
		switch r {
		case '(':
			total++
		case ')':
			total--
		}
	}

	return total
}

func part2(input string) interface{} {
	var c int
	var p int
	for i, r := range input {
		switch r {
		case '(':
			c++
		case ')':
			c--
		}

		if c == -1 {
			p = i + 1
			break
		}
	}

	return p
}

func ReadFileAsString(filePath string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(path + filePath)
	if err != nil {
		panic(err)
	}

	return string(dat)
}
