package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2024/day5/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

type conf struct {
	found    bool
	lessThan []string
}

func part1(input []string) interface{} {
	var i int
	config := make(map[string]conf)

	for {
		if input[i] == "" {
			i++
			break
		}

		parts := strings.Split(input[i], "|")
		config[parts[0]] = conf{found: false, lessThan: append(config[parts[0]].lessThan, parts[1])}

		i++
	}

	var sum int
	for {
		if i == len(input)-1 {
			break
		}

		tempConf := make(map[string]conf)
		for k, v := range config {
			tempConf[k] = v
		}

		valid := true
		vals := strings.Split(input[i], ",")
		for _, v := range vals {
			cf := tempConf[v]
			for _, lt := range cf.lessThan {
				if tempConf[lt].found {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
			tempConf[v] = conf{found: true, lessThan: cf.lessThan}
		}

		if valid {
			nr, _ := strconv.Atoi(vals[len(vals)/2])
			sum += nr
		}
		i++
	}

	return sum
}

var config map[string][]string

type ByConfig []string

func (a ByConfig) Len() int { return len(a) }
func (a ByConfig) Less(i, j int) bool {
	v := config[a[i]]
	for _, k := range v {
		if a[j] == k {
			return true
		}
	}

	return false
}
func (a ByConfig) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func part2(input []string) interface{} {
	config = make(map[string][]string)
	var i int
	for {
		if input[i] == "" {
			i++
			break
		}

		parts := strings.Split(input[i], "|")
		if v, ok := config[parts[0]]; ok {
			config[parts[0]] = append(v, parts[1])
		} else {
			config[parts[0]] = []string{parts[1]}
		}

		i++
	}

	var sum int
	for {
		if i == len(input) {
			break
		}

		vals := strings.Split(input[i], ",")
		if sort.IsSorted(ByConfig(vals)) {
			i++
			continue
		}

		sort.Sort(ByConfig(vals))
		nr, _ := strconv.Atoi(vals[len(vals)/2])
		sum += nr
		i++
	}

	return sum
}
