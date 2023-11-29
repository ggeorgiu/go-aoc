package main

import (
	"fmt"
	"strconv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2015/day7/part1.input")
	respPart1 := part1(input)
	fmt.Printf("P1: %v\n", respPart1)
	fmt.Printf("P2: %v\n", part2(input, respPart1))
}

func part1(input []string) int {
	circuit, config := buildConfig(input)
	step := true

	for step {
		step = false
		for k, v := range circuit {
			if val, ok := processInstruction(v, config); ok {
				config[k] = val
				step = true
				delete(circuit, k)
			}
		}
	}

	return config["a"]
}

func part2(input []string, respPart1 int) int {
	circuit, config := buildConfig(input)
	circuit["b"] = strconv.Itoa(respPart1)
	step := true

	for step {
		step = false
		for k, v := range circuit {
			if val, ok := processInstruction(v, config); ok {
				config[k] = val
				step = true
				delete(circuit, k)
			}
		}
	}

	return config["a"]
}

type operation struct {
	op1  int
	op2  int
	gate string
}

func processInstruction(v string, config map[string]int) (int, bool) {
	if isNumber(v) {
		return toNumber(v), true
	}

	if val, ok := config[v]; ok && val != -1 {
		return config[v], true
	}

	if !strings.Contains(v, " ") {
		return 0, false
	}

	op, shouldDo := buildOperation(v, config)
	if shouldDo {
		return doOperation(op), true
	}

	return 0, false
}

func buildConfig(input []string) (map[string]string, map[string]int) {
	var circuit = map[string]string{}
	var values = map[string]int{}

	for _, v := range input {
		split := strings.Split(v, " -> ")
		circuit[split[1]] = split[0]
		values[split[1]] = -1
	}

	return circuit, values
}

func evaluate(config map[string]int, k string) int {
	if isNumber(k) {
		return toNumber(k)
	}

	return config[k]
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func toNumber(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func doOperation(op operation) int {
	var v = 0
	switch op.gate {
	case "AND":
		v = op.op1 & op.op2
	case "OR":
		v = op.op1 | op.op2
	case "RSHIFT":
		v = op.op1 >> op.op2
	case "LSHIFT":
		v = op.op1 << op.op2
	default:
		v = ^op.op1
	}

	return int(uint16(v))
}

func buildOperation(i string, config map[string]int) (operation, bool) {
	split := strings.Split(i, " ")

	switch {
	case strings.Contains(i, "NOT"):
		if evaluate(config, split[1]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[0],
			op1:  evaluate(config, split[1]),
		}, true
	case strings.Contains(i, "AND") || strings.Contains(i, "OR"):
		if evaluate(config, split[0]) == -1 || evaluate(config, split[2]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[1],
			op1:  evaluate(config, split[0]),
			op2:  evaluate(config, split[2]),
		}, true
	default:
		if evaluate(config, split[0]) == -1 {
			return operation{}, false
		}
		return operation{
			gate: split[1],
			op1:  evaluate(config, split[0]),
			op2:  toNumber(split[2]),
		}, true
	}
}
