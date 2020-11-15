package search

import "testing"

func TestShortestPathWithBfs(t *testing.T) {
	adj := make([][]int, 7, 0)
	adj[0] = append(adj[0], 1)
	adj[0] = append(adj[0], 1)
	adj[0] = append(adj[0], 5)
	adj[1] = append(adj[0], 0)
	adj[1] = append(adj[1], 5)
	adj[2] = append(adj[2], 3)
	adj[3] = append(adj[3], 2)
	adj[3] = append(adj[3], 4)
	adj[3] = append(adj[3], 6)
	adj[4] = append(adj[4], 3)
	adj[5] = append(adj[5], 0)
	adj[5] = append(adj[5], 1)
	adj[6] = append(adj[6], 3)

	AssertEquals(t, 2, CountForest(adj))
}

func TestCheckFindCycle(t *testing.T) {
	adj := make([][]int, 4, 0)
	adj[0] = append(adj[0], 3)
	adj[0] = append(adj[0], 2)
	adj[1] = append(adj[1], 0)
	adj[2] = append(adj[2], 1)

	AssertTrue(t, HasCycle(adj))
}

func TestShouldNotFindCycle(t *testing.T) {
	adj := make([][]int, 4, 0)
	adj[0] = append(adj[0], 3)
	adj[0] = append(adj[0], 2)
	adj[1] = append(adj[1], 0)

	AssertFalse(t, HasCycle(adj))
}

func TestShouldFindCycleInStronglyConnectedGraph(t *testing.T) {
	adj := make([][]int, 2, 0)
	adj[0] = append(adj[0], 1)
	adj[1] = append(adj[1], 0)

	AssertTrue(t, HasCycle(adj))
}

func TestShouldGetTopologicalSort(t *testing.T) {
	adj := make([][]int, 3, 0)
	adj[0] = append(adj[0], 1)
	adj[0] = append(adj[0], 2)
	adj[1] = append(adj[1], 3)

	expected := []int{0, 1, 3, 2}

	AssertIntArrayEquals(t, expected, TopologicalSort(adj))
}

func TestShouldCheckCycleProperty(t *testing.T) {
	adj := make([][]int, 3, 0)
	adj[0] = append(adj[0], 1)
	adj[0] = append(adj[0], 2)
	adj[1] = append(adj[1], 3)

	AssertFalse(t, HasCycle(adj))

	adj2 := make([][]int, 4, 0)
	adj[0] = append(adj[0], 1)
	adj[1] = append(adj[1], 3)
	adj[1] = append(adj[1], 2)
	adj[2] = append(adj[2], 0)

	AssertTrue(t, HasCycle(adj2))
}
