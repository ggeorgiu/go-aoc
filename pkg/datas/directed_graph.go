package datas

import (
	"fmt"
	"strings"
)

type DirectedGraph struct {
	Nodes map[Node][]Edge
}

type Node string

func (n *Node) String() string {
	return string(*n)
}

type Edge struct {
	Weight int
	Node   Node
}

func (g *DirectedGraph) AddNode(n Node) {
	if g.Nodes == nil {
		g.Nodes = make(map[Node][]Edge)
	}

	g.Nodes[n] = []Edge{}
}

func (g *DirectedGraph) AddEdge(from string, to string, weight int) {
	vFrom := Node(from)
	vTo := Node(to)

	if _, ok := g.Nodes[vFrom]; !ok {
		g.AddNode(vFrom)
	}

	if _, ok := g.Nodes[vTo]; !ok {
		g.AddNode(vTo)
	}

	g.Nodes[vFrom] = append(g.Nodes[vFrom], Edge{
		Weight: weight,
		Node:   vTo,
	})
	g.Nodes[vTo] = append(g.Nodes[vTo], Edge{
		Weight: weight,
		Node:   vFrom,
	})
}

func (g *DirectedGraph) GetNode(key string) Node {
	for k := range g.Nodes {
		if k.String() == key {
			return k
		}
	}

	panic("node not found")
}

func (g *DirectedGraph) GetEdges(key string) []Edge {
	for k, v := range g.Nodes {
		if k.String() == key {
			return v
		}
	}

	panic("node not found")
}

func (g *DirectedGraph) String() string {
	var sb strings.Builder
	for k, v := range g.Nodes {
		sb.WriteString(fmt.Sprintf("%s: %v\n", k, v))
	}

	return sb.String()
}
