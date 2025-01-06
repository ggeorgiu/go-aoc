package main

import (
	"fmt"
	"strconv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsString("/2024/day9/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

func part1(input string) interface{} {
	var isFile = true
	var data []int
	var id int

	for _, v := range input {
		nr, _ := strconv.Atoi(string(v))
		for j := 0; j < nr; j++ {
			if isFile {
				data = append(data, id)
			} else {
				data = append(data, -1)
			}
		}

		if isFile {
			id++
		}
		isFile = !isFile
	}

	ids := 0
	ide := len(data) - 1

	for ids < ide {
		for data[ids] != -1 && ids < ide {
			ids++
		}
		for data[ide] == -1 && ids < ide {
			ide--
		}

		data[ids] = data[ide]
		data[ide] = -1
	}

	ids = 0
	checksum := 0

	for data[ids] != -1 {
		checksum += ids * data[ids]
		ids++
	}

	return checksum
}

func part2(input string) interface{} {
	var isFile = true
	var data []int
	var id int

	for _, v := range input {
		nr, _ := strconv.Atoi(string(v))
		for j := 0; j < nr; j++ {
			if isFile {
				data = append(data, id)
			} else {
				data = append(data, -1)
			}
		}

		if isFile {
			id++
		}
		isFile = !isFile
	}

	ids := 0
	ide := len(data) - 1

	for ide > 0 {
		fileIdx := ide
		for data[fileIdx] == -1 {
			fileIdx--
		}

		lenFile := 1
		for ids < fileIdx && data[fileIdx] != -1 && data[fileIdx] == data[fileIdx-1] {
			fileIdx--
			lenFile++
		}

		found := false
		before := true
		spaceIdx := 0
		lenSpace := 0

		for !found && before {
			for data[spaceIdx] != -1 && spaceIdx < fileIdx {
				spaceIdx++
			}

			tempIds := spaceIdx
			for data[tempIds] == -1 {
				if tempIds >= fileIdx {
					before = false
					break
				}
				tempIds++
				lenSpace++
			}

			if lenSpace >= lenFile {
				found = true
			} else {
				lenSpace = 0
				spaceIdx++
			}

			if found {
				for i := 0; i < lenFile; i++ {
					data[spaceIdx+i] = data[fileIdx+i]
					data[fileIdx+i] = -1
				}
			}

			ide = fileIdx - 1
		}
	}

	checksum := 0

	for i, v := range data {
		if v == -1 {
			continue
		}

		checksum += i * v
	}

	return checksum
}
