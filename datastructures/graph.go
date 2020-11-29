package datastructures

import (
	"galgo/algorithms/graph/mst"
	"math"
)

type Graph struct {
	isDirected bool
	edges      []Edge
	adjMatrix  [][]int
	adjList    [][]Pair
	n          int
}

type Edge struct {
	X *graphnode
	Y *graphnode
	W int
}

func (e *Edge) Compare(x *Edge) int {
	if e.W < x.W {
		return -1
	} else {
		return 1
	}
}

func newUnweightedEdge(x *graphnode, y *graphnode) Edge {
	return Edge{x, y, 1}
}

func NewGraph(n int) Graph {
	return Graph{false, make([]Edge, 0), make([][]int, 0, 0), make([][]Pair, 0, 0), n}
}

func (g *Graph) buildAdjacencyMatrix() {
	g.adjMatrix = make([][]int, g.n, g.n)

	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			g.adjMatrix[i][j] = math.MaxInt64
		}
	}

	for i := 0; i < len(g.edges); i++ {
		g.adjMatrix[g.edges[i].X.Value][g.edges[i].Y.Value] = g.edges[i].W
		if !g.isDirected {
			g.adjMatrix[g.edges[i].Y.Value][g.edges[i].X.Value] = g.edges[i].W
		}
	}
}

func (g *Graph) buildAdjacencyList() {
	g.adjList = make([][]Pair, g.n, 0)

	for i := 0; i < len(g.edges); i++ {
		g.adjList[g.edges[i].X.Value] = append(g.adjList[g.edges[i].X.Value], Pair{g.edges[i].Y.Value, g.edges[i].W})
	}
}

func (g *Graph) GetAdjacencyMatrix() [][]int {
	if g.adjMatrix == nil {
		g.buildAdjacencyMatrix()
	}

	return g.adjMatrix
}

func (g *Graph) GetAdjacencyList() [][]Pair {
	if g.adjList == nil {
		g.buildAdjacencyList()
	}

	return g.adjList
}

func (g *Graph) GetEdges() []Edge {
	return g.edges
}

func (g *Graph) Mst() []Edge {
	return mst.Kruskal(g.edges, g.n)
}

func (g *Graph) AddEdges(edges []Edge) {
	for i := 0; i < len(edges); i++ {
		g.AddEdge(edges[i])
	}
}

func (g *Graph) AddEdge(edge Edge) {
	g.edges = append(g.edges, edge)

	if !g.isDirected {
		nodeX := newgraphnode(edge.Y.Value)
		nodeY := newgraphnode(edge.X.Value)
		g.edges = append(g.edges, Edge{&nodeX, &nodeY, edge.W})
	}
}

func (g *Graph) AddWeightEdge(x int, y int, w int) {
	nodeX := newgraphnode(x)
	nodeY := newgraphnode(y)
	g.AddEdge(Edge{&nodeX, &nodeY, w})
}

type graphnode struct {
	Value int
}

func newgraphnode(v int) graphnode {
	return graphnode{v}
}
