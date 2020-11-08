package search

import "testing"

func TestFindIndex(t *testing.T) {
	a := []int {1, 2, 3, 4, 5, 6}
	AssertEquals(t, BinarySearch(a, 1), 0)
	AssertEquals(t, BinarySearch(a, 2), 1)
	AssertEquals(t, BinarySearch(a, 4), 3)
	AssertEquals(t, BinarySearch(a, -1), -1)
	AssertEquals(t, BinarySearch(a, 7), -1)
	AssertEquals(t, BinarySearch(make([]int, 0), 7), -1)
}

func TestFindIndexFirstLowest(t *testing.T) {
	a := []int {1, 2, 3, 4, 5, 6}
	AssertEquals(t, 1, BinarySearchFirstLowest(a, 2))
	AssertEquals(t, -1, BinarySearchFirstLowest(a, 0))
	AssertEquals(t, 2, BinarySearchFirstLowest(a, 4))
	AssertEquals(t, 5, BinarySearchFirstLowest(a, 8))
}

func TestFindIndexFirstHighest(t *testing.T) {
	a := []int {1, 2, 3, 4, 5, 6}
	AssertEquals(t, 1, BinarySearchFirstHighest(a, 2))
	AssertEquals(t, 0, BinarySearchFirstHighest(a, 0))
	AssertEquals(t, 3, BinarySearchFirstHighest(a, 4))
	AssertEquals(t, 6, BinarySearchFirstHighest(a, 8))
}

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func AssertTrue(t *testing.T, actual bool)  {
	if !actual {
		t.Errorf("Expected %v, actual %v", true, actual)
	}
}

func AssertFalse(t *testing.T, actual bool)  {
	if actual {
		t.Errorf("Expected %v, actual %v", false, actual)
	}
}