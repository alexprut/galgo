package sorting

import (
	"reflect"
	"testing"
)

func TestShouldSort(t *testing.T) {
	var toSort = []int {5, 2, 1, 9, 0, 33, 3, 3, 0}
	var expected = []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	AssertIntArrayEquals(t, expected, BubbleSort(toSort))
}

func TestSortBubbleSort(t *testing.T) {
	toSort := []int {5, 2, 1, 9, 0, 33, 3, 3, 0}
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	AssertIntArrayEquals(t, expected, BubbleSort(toSort))
}

func TestSortInverseSortedBubbleSort(t *testing.T) {
	toSort := []int {33, 9, 5, 3, 3, 2, 1, 0, 0}
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	AssertIntArrayEquals(t, expected, BubbleSort(toSort))
}

func TestSortSortedBubbleSort(t *testing.T) {
	expected := []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	AssertIntArrayEquals(t, expected, BubbleSort(expected))
}

func TestHandleEdgeCasesBubbleSort(t *testing.T) {
	AssertIntArrayEquals(t, []int {}, BubbleSort([]int {}))
	AssertIntArrayEquals(t, []int {1}, BubbleSort([]int {1}))
}

func AssertIntArrayEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}