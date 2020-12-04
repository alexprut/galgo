package datastructures

import (
	"math"
	"reflect"
	"testing"
)

func TestShouldCreateUndirectedGraph(t *testing.T) {
	n := 4
	graph := NewGraph(n)
	graph.AddWeightEdge(0, 1, 1)
	graph.AddWeightEdge(0, 2, 1)

	expected := make([][]int, n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			expected[i][j] = math.MaxInt64
		}
	}
	expected[0][1] = 1
	expected[1][0] = 1
	expected[0][2] = 1
	expected[2][0] = 1

	AssertIntIntArrayEquals(t, expected, graph.GetAdjacencyMatrix())
}

func TestShouldCreateDirectedGraph(t *testing.T) {
	n := 4
	graph := NewGraph(n)
	graph.isDirected = true
	graph.AddWeightEdge(0, 1, 1)
	graph.AddWeightEdge(0, 2, 1)

	expected := make([][]int, n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			expected[i][j] = math.MaxInt64
		}
	}
	expected[0][1] = 1
	expected[0][2] = 1

	AssertIntIntArrayEquals(t, expected, graph.GetAdjacencyMatrix())
}

func AssertIntArrayEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func AssertIntIntArrayEquals(t *testing.T, expected [][]int, actual [][]int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
