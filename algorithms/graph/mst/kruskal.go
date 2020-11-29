package mst

import (
	"galgo/datastructures"
	"sort"
)

func Kruskal(edges []datastructures.Edge, n int) []datastructures.Edge {
	mst := make([]datastructures.Edge, 0)
	sets := datastructures.NewDisjointSet()

	for i := 0; i < n; i++ {
		sets.MakeSet(i)
	}

	sort.SliceStable(edges, func(i, j int) bool { return datastructures.Less(edges[i], edges[j]) })

	for i := 0; i < len(edges); i++ {
		if sets.FindSet(edges[i].X.Value) != sets.FindSet(edges[i].Y.Value) {
			sets.Union(sets.GetElement(edges[i].X.Value), sets.GetElement(edges[i].Y.Value))
			mst = append(mst, edges[i])
		}
	}

	return mst
}
