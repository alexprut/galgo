package graph

import (
	"galgo/datastructures"
	"math"
)

func BellmanFord(edges []datastructures.Edge, n int, start int) []int {
	distance := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt64
		parent[i] = -1
	}
	distance[start] = 0
	for i := 0; i < n; i++ {
		for e := 0; e < len(edges); e++ {
			if distance[edges[e].Y.Value] > distance[edges[e].X.Value]+edges[e].W {
				distance[edges[e].Y.Value] = distance[edges[e].X.Value] + edges[e].W
				parent[edges[e].Y.Value] = edges[e].X.Value
			}
		}
	}

	for e := 0; e < len(edges); e++ {
		if distance[edges[e].Y.Value] > distance[edges[e].X.Value]+edges[e].W {
			panic("Graph contains a negative cycle")
		}
	}

	return distance
}
