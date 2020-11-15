package algorithms

import "testing"

func TestShouldFindMaxSubarray(t *testing.T) {
	AssertEquals(t, 0, MaximumSubarray([]int{}))
	AssertEquals(t, 6, MaximumSubarray([]int{1, -1, -2, 4, -1, 3}))
}
