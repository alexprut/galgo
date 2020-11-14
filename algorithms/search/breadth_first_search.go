package search

import (
	"galgo/datastructures"
	"math"
)

func ShortestPath(adj [][]int, start int) []int {
	var shortestPaths = make([]int, len(adj))
	var isVisited = make([]bool, len(adj))
	for i := 0; i < len(shortestPaths); i++ {
		shortestPaths[i] = math.MaxInt64
	}
	var queue = datastructures.NewQueue()
	queue.Enqueue(start)
	isVisited[start] = true
	shortestPaths[start] = 0

	for !queue.Empty() {
		var edge = queue.Dequeue().(int)
		for i := 0; i < len(adj[edge]); i++ {
			y := adj[edge][i]
			if !isVisited[y] {
				isVisited[y] = true
				shortestPaths[y] = shortestPaths[edge] + 1
				queue.Enqueue(y)
			}
		}
	}

	return shortestPaths
}

func PredecessorSubgraph(adj [][]int, start int) []int {
	parent := make([]int, len(adj))
	isVisited := make([]bool, len(adj))
	for i := 0; i < len(parent); i++ {
		parent[i] = math.MaxInt64
	}
	queue := datastructures.NewQueue()
	queue.Enqueue(start)
	isVisited[start] = true
	for !queue.Empty() {
		edge := queue.Dequeue().(int)
		for i := 0; i < len(adj[edge]); i++ {
			y := adj[edge][i]
			if !isVisited[y] {
				isVisited[y] = true
				parent[y] = edge
				queue.Enqueue(y)
			}
		}
	}

	return parent
}
