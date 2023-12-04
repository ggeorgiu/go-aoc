package main

import (
	"fmt"
	"go-aoc/pkg/gen"
	"math"
	"sort"
	"strconv"
	"strings"

	"go-aoc/pkg/datas"

	"go-aoc/pkg/in"
)

func main() {
	input := in.ReadFileAsStringSlice("/2015/day9/part1.input")
	fmt.Printf("P1: %v\n", part1(input))
	fmt.Printf("P1: %v\n", part2(input))
}

func part1(input []string) interface{} {
	g := datas.DirectedGraph{}
	var nodes []string
	var ids []int
	for _, l := range input {
		from, to, val := parse(l)
		g.AddEdge(from, to, val)
	}

	i := 0
	for k := range g.Nodes {
		nodes = append(nodes, k.String())
		ids = append(ids, i)
		i++
	}
	sort.Strings(nodes)
	perms := gen.Permutations(ids)

	m := math.MaxInt
	for _, v := range perms {
		cost := 0
		current := g.GetNode(nodes[v[0]])

		for k := 0; k < len(v); k++ {
			edges := g.GetEdges(current.String())
			for _, e := range edges {
				if e.Node.String() == nodes[v[k]] {
					cost += e.Weight
					current = e.Node
					break
				}
			}
		}

		if cost < m {
			m = cost
		}
	}

	return m
}

func part2(input []string) interface{} {
	g := datas.DirectedGraph{}
	var nodes []string
	var ids []int
	for _, l := range input {
		from, to, val := parse(l)
		g.AddEdge(from, to, val)
	}

	i := 0
	for k := range g.Nodes {
		nodes = append(nodes, k.String())
		ids = append(ids, i)
		i++
	}
	sort.Strings(nodes)

	perms := gen.Permutations(ids)
	m := math.MinInt

	for _, v := range perms {
		cost := 0
		current := g.GetNode(nodes[v[0]])

		for k := 0; k < len(v); k++ {
			edges := g.GetEdges(current.String())

			for _, e := range edges {
				if e.Node.String() == nodes[v[k]] {
					cost += e.Weight
					current = e.Node
					break
				}
			}
		}

		if cost > m {
			m = cost
		}
	}

	return m
}

func parse(v string) (string, string, int) {
	split := strings.Split(v, " = ")
	dist, _ := strconv.Atoi(split[1])

	nodes := strings.Split(split[0], " to ")

	return nodes[0], nodes[1], dist
}
