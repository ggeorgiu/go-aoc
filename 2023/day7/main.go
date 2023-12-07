package main

import (
	"fmt"
	"sort"
	"strings"

	"go-aoc/pkg/conv"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2023/day7/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P2: %v\n", part2(input))
}

var conf map[uint8]int

type hand struct {
	cards string
	rank  int
	bid   int
}

type hands []hand

func (h hands) Len() int {
	return len(h)
}

func (h hands) Less(i, j int) bool {
	one := h[i]
	other := h[j]

	if one.rank != other.rank {
		return one.rank < other.rank
	}

	var k int
	for k = 0; k < len(one.cards); k++ {
		if one.cards[k] != other.cards[k] {
			break
		}
	}

	return conf[one.cards[k]] < conf[other.cards[k]]
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func part1(input []string) interface{} {
	conf = make(map[uint8]int)
	for i, v := range "23456789TJQK" {
		conf[uint8(v)] = i
	}

	var hs hands
	for _, line := range input {
		parts := strings.Split(line, " ")
		cards := parts[0]
		bid := conv.ToInt(parts[1])
		rk := getRank(cards)

		hs = append(hs, hand{
			cards: cards,
			rank:  rk,
			bid:   bid,
		})
	}

	sort.Sort(hs)

	var sum int
	for i, v := range hs {
		sum += v.bid * (i + 1)
	}

	return sum
}

func part2(input []string) interface{} {
	conf = make(map[uint8]int)
	for i, v := range "J23456789TQK" {
		conf[uint8(v)] = i
	}

	var hs hands
	for _, line := range input {
		parts := strings.Split(line, " ")
		cards := parts[0]
		bid := conv.ToInt(parts[1])

		mxRank := -1
		for k := range conf {
			if k == 'J' {
				continue
			}
			auxCards := cards
			auxCards = strings.ReplaceAll(auxCards, "J", string(k))

			rk := getRank(auxCards)
			if rk > mxRank {
				mxRank = rk
			}
		}

		hs = append(hs, hand{
			cards: cards,
			rank:  mxRank,
			bid:   bid,
		})
	}

	sort.Sort(hs)

	var sum int
	for i, v := range hs {
		sum += v.bid * (i + 1)
	}

	return sum
}

// getRank this one can really be improved
func getRank(cards string) int {
	cardCount := make(map[rune]int)
	for _, c := range cards {
		cardCount[c]++
	}

	switch len(cardCount) {
	case 1:
		return 6
	case 2:
		for _, count := range cardCount {
			if count == 1 || count == 4 {
				return 5
			}
			return 4
		}
	case 3:
		mxCount := -1
		for _, count := range cardCount {
			if count > mxCount {
				mxCount = count
			}
		}

		if mxCount == 3 {
			return 3
		}

		return 2
	case 4:
		return 1
	}

	return 0
}
