package search

import (
	"math"
	"reflect"
	"testing"
)

func TestComputeShortestPathWithBfs(t *testing.T) {
	adj := make([][]int, 0, 0)
	adj[0] = append(adj[0], 1)

	adj[1] = append(adj[1], 0)
	adj[1] = append(adj[1], 2)
	adj[2] = append(adj[2], 1)
	adj[2] = append(adj[2], 3)
	adj[3] = append(adj[3], 2)
	adj[3] = append(adj[3], 5)
	adj[3] = append(adj[3], 6)
	adj[3] = append(adj[3], 4)
	adj[4] = append(adj[4], 3)
	adj[5] = append(adj[5], 3)
	adj[6] = append(adj[6], 3)

	expectedShortestPath := []int{0, 1, 2, 3, 4, 4, 4}
	shortestPath := ShortestPath(adj, 0)
	AssertIntArrayEquals(t, expectedShortestPath, shortestPath)
}

func TestComputePredecessorSubGraph(t *testing.T) {
	adj := make([][]int, 0, 0)
	adj[0] = append(adj[0], 1)
	adj[1] = append(adj[1], 0)
	adj[1] = append(adj[1], 2)
	adj[2] = append(adj[2], 1)
	adj[2] = append(adj[2], 3)
	adj[3] = append(adj[3], 2)
	adj[3] = append(adj[3], 5)
	adj[3] = append(adj[3], 6)
	adj[3] = append(adj[3], 4)
	adj[4] = append(adj[4], 3)
	adj[4] = append(adj[4], 6)
	adj[5] = append(adj[5], 3)
	adj[5] = append(adj[5], 6)
	adj[6] = append(adj[6], 3)
	adj[6] = append(adj[6], 4)
	adj[6] = append(adj[6], 5)

	expectedShortestPath := []int{math.MaxInt64, 0, 1, 2, 3, 3, 3}
	shortestPath := PredecessorSubgraph(adj, 0)
	AssertIntArrayEquals(t, expectedShortestPath, shortestPath)
}

func AssertIntArrayEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
