package main

import (
	"fmt"
	"go-aoc/pkg/conv"
	"strings"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2016/day10/part1.input")
	fmt.Printf("P1/P2: %v\n", part(input))
}

type bot struct {
	value1 int
	value2 int

	highID string
	lowID  string
}

func (b *bot) receive(v int) {
	if b.value1 == 0 {
		b.value1 = v
		return
	}

	if v < b.value1 {
		b.value2 = b.value1
		b.value1 = v
		return
	}

	b.value2 = v
}

func part(input []string) any {
	var bots = make(map[string]*bot)
	var outputs = make(map[string]int)

	for _, line := range input {
		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "value") {
			id := parts[5]
			value := parts[1]

			var b *bot
			b, ok := bots[id]
			if !ok {
				b = &bot{}
				bots[id] = b
			}

			b.receive(conv.ToInt(value))
			continue
		}

		sourceID, lowID, highID := parts[1], parts[6], parts[11]
		destLow, destHigh := parts[5], parts[10]

		var b *bot
		b, ok := bots[sourceID]
		if !ok {
			b = &bot{}
			bots[sourceID] = b
		}

		b.lowID = lowID
		if destLow == "output" {
			b.lowID = "o-" + lowID
		}

		b.highID = highID
		if destHigh == "output" {
			b.highID = "o-" + highID
		}

	}

	var proceeded = true
	var targetNode string
	lowTarget, highTarget := 17, 61

	for proceeded {
		proceeded = false
		for k, v := range bots {
			if v.value1 == 0 || v.value2 == 0 {
				continue
			}

			if v.value1 == lowTarget && v.value2 == highTarget {
				targetNode = k
			}

			if strings.HasPrefix(v.lowID, "o-") {
				outputs[v.lowID[2:]] = v.value1
			} else {
				low, _ := bots[v.lowID]
				low.receive(v.value1)
			}
			v.value1 = 0

			if strings.HasPrefix(v.highID, "o-") {
				outputs[v.highID[2:]] = v.value2
			} else {
				high, _ := bots[v.highID]
				high.receive(v.value2)
			}
			v.value2 = 0

			proceeded = true
		}
	}

	return fmt.Sprintf("%s / %d", targetNode, outputs["0"]*outputs["1"]*outputs["2"])
}
