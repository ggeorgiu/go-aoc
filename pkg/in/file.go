package in

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileAsString(filePath string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(path + filePath)
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(string(dat), "\n")
}

func ReadFileAsStringSlice(filePath string) []string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path + filePath)
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileAsRuneMatrix(filePath string) [][]rune {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path + filePath)
	if err != nil {
		panic(err)
	}

	var m [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, r := range line {
			row[i] = r
		}
		m = append(m, row)
	}

	return m
}
