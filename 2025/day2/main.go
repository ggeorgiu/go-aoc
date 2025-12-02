package main

import (
	"fmt"
	"go-aoc/pkg/conv"
	"go-aoc/pkg/in"
	"math"
	"strings"
)

func main() {
	input := in.ReadFileAsString("/2025/day2/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input string) any {
	intervals := strings.Split(input, ",")

	var count int
	for _, itv := range intervals {
		bounds := strings.Split(itv, "-")
		start, stop := conv.ToInt(bounds[0]), conv.ToInt(bounds[1])

		for i := start; i <= stop; i++ {
			if isInvalid(i) {
				count += i
			}
		}
	}

	return count
}

func isInvalid(nr int) bool {
	var l int

	aux := nr
	for aux > 0 {
		aux = aux / 10
		l++
	}

	if l%2 != 0 {
		return false
	}

	exp := int(math.Pow10(l / 2))
	return (nr / exp) == (nr % exp)
}

func part2(input string) any {
	intervals := strings.Split(input, ",")

	var count int
	for _, itv := range intervals {
		bounds := strings.Split(itv, "-")
		start, stop := conv.ToInt(bounds[0]), conv.ToInt(bounds[1])

		for i := start; i <= stop; i++ {
			if isInvalid2(i) {
				count += i
			}
		}
	}

	return count
}

func isInvalid2(nr int) bool {
	if nr < 10 {
		return false
	}

	var l int
	var digits []int

	aux := nr
	for aux > 0 {
		digits = append(digits, aux%10)
		aux = aux / 10
		l++
	}
	var valid bool

	if l%2 == 0 {
		half := l / 2
		for i := 1; i <= half; i++ {
			if l%i != 0 {
				continue
			}

			valid = true
			k := len(digits) - 1
			for k-i >= 0 {
				if digits[k] != digits[k-i] {
					valid = false
					break
				}
				k--
			}
			if valid {
				return true
			}
		}
	} else {
		rad := int(math.Sqrt(float64(l)))
		for i := 1; i <= rad; i += 2 {
			valid = true
			for k := len(digits) - 1; k-i >= 0; k-- {
				if digits[k] != digits[k-i] {
					valid = false
					break
				}
			}
			if valid {
				return true
			}
		}
	}

	return false
}
