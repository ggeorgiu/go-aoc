package main

import (
	"fmt"
	"go-aoc/pkg/conv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2025/day6/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	input2 := in.ReadFileAsRuneMatrix("/2025/day6/part1.input")
	fmt.Printf("P1: %v\n", part2(input2))
}

type problem struct {
	numbers []int
	sign    string
}

func (p problem) solve() int {
	op := func(a int, b int) int {
		if p.sign == "+" {
			return a + b
		}

		return a * b
	}

	var total int
	if p.sign == "*" {
		total = 1
	}

	for _, v := range p.numbers {
		total = op(total, v)
	}

	return total
}

func part1(input []string) any {
	problems := make([]problem, len(strings.Fields(input[0])))

	for _, line := range input {
		parts := strings.Fields(line)
		for i, part := range parts {
			prob := problems[i]
			if part == "+" || part == "*" {
				prob.sign = part
				problems[i] = prob
				continue
			}

			prob.numbers = append(prob.numbers, conv.ToInt(part))
			problems[i] = prob
		}
	}

	var total int
	for _, p := range problems {
		total += p.solve()
	}

	return total
}

func part2(input [][]rune) any {
	var problems []problem
	var pb problem
	for j := len(input[0]) - 1; j >= 0; j-- {
		sb := strings.Builder{}
		for i := 0; i < len(input); i++ {
			if input[i][j] == '+' || input[i][j] == '*' {
				pb.numbers = append(pb.numbers, conv.ToInt(strings.Fields(sb.String())[0]))
				pb.sign = string(input[i][j])
				problems = append(problems, pb)

				pb = problem{}
				sb = strings.Builder{}

				break
			}

			sb.WriteRune(input[i][j])
		}
		nr := sb.String()
		fields := strings.Fields(nr)
		if len(fields) == 0 {
			continue
		}
		pb.numbers = append(pb.numbers, conv.ToInt(fields[0]))
	}

	var total int
	for _, p := range problems {
		total += p.solve()
	}

	return total
}
