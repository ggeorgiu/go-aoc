package in

import (
	"bufio"
	"os"
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

	return string(dat)
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
