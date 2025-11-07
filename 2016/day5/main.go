package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "wtnhxymk"
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input string) any {
	var count int
	var i int
	var res strings.Builder

	for count < 8 {
		result := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hexa := hex.EncodeToString(result[:])

		if hexa[0:5] == "00000" {
			count++
			res.WriteString(string(hexa[5]))
		}

		i++
	}

	return res.String()
}

func part2(input string) any {
	var i int
	res := make([]uint8, 8)
	done := make([]bool, 8)

	doneFn := func() bool {
		for _, v := range done {
			if !v {
				return false
			}
		}
		return true
	}

	for !doneFn() {
		result := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hexa := hex.EncodeToString(result[:])

		if hexa[0:5] == "00000" {
			pos := int(hexa[5] - '0')
			if pos < 0 || pos > 7 {
				i++
				continue
			}
			if done[pos] {
				i++
				continue
			}

			res[pos] = hexa[6]
			done[pos] = true
		}
		printResult(res, i)
		i++
	}

	return ""
}

func printResult(result []uint8, index int) {
	fmt.Print(fmt.Sprintf("\r%s - %d", string(result[:]), index))
	os.Stdout.Sync()
}
