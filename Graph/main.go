package main

import (
	"fmt"

	dfs "github.com/imrushi/algorithms/Graph/DFS"
	"github.com/imrushi/algorithms/Graph/graph"
)

func main() {
	g := graph.NewDirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)
	fmt.Println(g)
	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)
	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	dfs.DFS(g, g.Vertices[1], cb)
	fmt.Println(visitedOrder)
}
