package datastructures

import "testing"

func TestShouldUnionSets(t *testing.T) {
	sets := NewDisjointSet()
	a := sets.MakeSet(1)
	b := sets.MakeSet(2)
	sets.Union(a, b)

	AssertEquals(t, a.parent.value, b.value)
	AssertEquals(t, b.parent.value, b.value)
	AssertEquals(t, b.rank, 1)
}
