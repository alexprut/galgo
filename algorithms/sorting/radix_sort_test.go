package sorting

import "testing"

func TestRadixSort(t *testing.T) {
	toSort := []int {5, 2, 1, 9, 0, 33, 3, 3, 0}
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	RadixSort(toSort, 2)

	AssertIntArrayEquals(t, expected, toSort)
}

func TestShouldSortInverseSorted(t *testing.T) {
	toSort := []int {33, 9, 5, 3, 3, 2, 1, 0, 0}
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	RadixSort(toSort, 2)

	AssertIntArrayEquals(t, expected, toSort)
}

func TestShouldSortSorted(t *testing.T) {
	toSort := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	RadixSort(toSort, 2)

	AssertIntArrayEquals(t, expected, toSort)
}

func TestShouldHandleEdgeCases(t *testing.T) {
	empty := []int {}
	RadixSort(empty, 1)
	single := []int {1}
	RadixSort(single, 1)

	AssertIntArrayEquals(t, []int {}, empty)
	AssertIntArrayEquals(t, []int {1}, single)
}
