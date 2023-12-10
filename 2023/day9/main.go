package main

import (
	"fmt"
	"go-aoc/pkg/algs"
	"go-aoc/pkg/conv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day9/part1.input")
	//fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input []string) interface{} {
	var sum int
	for _, line := range input {
		is := conv.ToIntSlice(strings.Split(line, " "))
		sum += predict(is)
	}

	return sum
}

func predict(data []int) int {
	var lastNums []int
	var conf = data
	for {
		if allZeroes(conf) {
			break
		}
		fmt.Println(conf)

		conf2 := make([]int, len(conf)-1)
		for i := 0; i < len(conf)-1; i++ {
			conf2[i] = conf[i+1] - conf[i]
		}

		lastNums = append(lastNums, conf2[len(conf2)-1])
		conf = conf2
	}

	return data[len(data)-1] + algs.Sum(lastNums)
}

func part2(input []string) interface{} {
	var sum int
	for _, line := range input {
		is := conv.ToIntSlice(strings.Split(line, " "))
		p := predict2(is)
		sum += p
	}

	return sum
}

func predict2(data []int) int {
	var firstNums []int
	var conf = data
	for {
		if allZeroes(conf) {
			break
		}

		conf2 := make([]int, len(conf)-1)
		for i := 0; i < len(conf)-1; i++ {
			conf2[i] = conf[i+1] - conf[i]
		}

		firstNums = append(firstNums, conf2[0])
		conf = conf2
	}

	var pred int
	for i := len(firstNums) - 1; i >= 0; i-- {
		pred = firstNums[i] - pred
	}

	return data[0] - pred
}

func allZeroes(data []int) bool {
	for _, v := range data {
		if v != 0 {
			return false
		}
	}

	return true
}
