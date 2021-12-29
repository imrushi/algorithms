package dfs

import (
	"fmt"

	graph "github.com/imrushi/algorithms/Graph/graph"
)

func DFS(g *graph.Graph, startVertex *graph.Vertex, visitCb func(int)) {
	fmt.Println("\nDFS")
	visited := map[int]bool{}

	if startVertex == nil {
		return
	}
	fmt.Printf("visited: %v\n", visited)
	visited[startVertex.Key] = true

	visitCb(startVertex.Key)

	for _, v := range startVertex.Vertices {
		fmt.Printf("Visiting Adjustent: %v", v)
		if visited[v.Key] {
			continue
		}
		DFS(g, v, visitCb)
	}
}
