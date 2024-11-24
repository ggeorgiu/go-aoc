package in

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileAsString(filePath string) string {
	file := openFile(filePath)

	dat, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return strings.TrimSuffix(string(dat), "\n")
}

func ReadFileAsStringSlice(filePath string) []string {
	file := openFile(filePath)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileAsCSVLine(filePath string) []string {
	file := openFile(filePath)

	scanner := bufio.NewScanner(file)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	return strings.Split(line, ", ")
}

func ReadFileAsRuneMatrix(filePath string) [][]rune {
	file := openFile(filePath)

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

func openFile(filePath string) *os.File {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path + filePath)
	if err != nil {
		panic(err)
	}
	return file
}
