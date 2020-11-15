package search

import "galgo/datastructures"

func CountForest(adj [][]int) int {
	visited := make([]bool, len(adj))
	counter := 0
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			counter++
			dfsCountForest(adj, i, visited)
		}
	}
	return counter
}

func dfsCountForest(adj [][]int, start int, visited []bool) {
	stack := datastructures.NewStack()
	stack.Push(start)
	visited[start] = true
	for !stack.Empty() {
		node := stack.Pop().(int)
		for i := 0; i < len(adj[node]); i++ {
			y := adj[node][i]
			if !visited[y] {
				visited[y] = true
				stack.Push(y)
			}
		}
	}
}

func HasCycle(adj [][]int) bool {
	visited := make([]bool, len(adj))
	epoch := make([]int, len(adj))
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			hasCycle := dfsHasCycle(adj, i, visited, epoch)
			if hasCycle {
				return true
			}
		}
		return false
	}
	return false
}

func dfsHasCycle(adj [][]int, start int, visited []bool, epoch []int) bool {
	stack := datastructures.NewStack()
	stack.Push(start)
	for !stack.Empty() {
		node := stack.Pop().(int)
		visited[node] = true
		epoch[node] = start
		for i := 0; i < len(adj[node]); i++ {
			y := adj[node][i]
			if visited[y] && epoch[y] == start {
				return true
			}
			if !visited[y] {
				stack.Push(y)
			}
		}
	}
	return false
}

func IsDAG(adj [][]int) bool {
	return !HasCycle(adj)
}

func TopologicalSort(adj [][]int) []int {
	if !IsDAG(adj) {
		panic("Graph is not a DAG")
	}

	visited := make([]bool, len(adj))
	topologicalSort := make([]int, 0)
	for i := 0; i < len(adj); i++ {
		if !visited[i] {
			DfsTopologicalSort(adj, i, visited, topologicalSort)
		}
	}
	return topologicalSort
}

func DfsTopologicalSort(adj [][]int, node int, visited []bool, topologicalSort []int) {
	for i := 0; i < len(adj[node]); i++ {
		y := adj[node][i]
		if !visited[y] {
			DfsTopologicalSort(adj, y, visited, topologicalSort)
		}
	}
	topologicalSort = append(topologicalSort, node)
	visited[node] = true
}
