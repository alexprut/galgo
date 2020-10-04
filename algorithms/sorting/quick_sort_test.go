package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var toSort = []int{5, 2, 1, 9, 0, 33, 3, 3, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	QuickSort(toSort, 0, len(toSort))

	if !reflect.DeepEqual(expected, toSort) {
		t.Errorf("Expected %v, actual %v", expected, toSort)
	}
}

func TestQuickSortInverseSorted(t *testing.T) {
	var toSort = []int{33, 9, 5, 3, 3, 2, 1, 0, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	QuickSort(toSort, 0, len(toSort))

	if !reflect.DeepEqual(expected, toSort) {
		t.Errorf("Expected %v, actual %v", expected, toSort)
	}
}

func TestQuickSortSorted(t *testing.T) {
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	QuickSort(expected, 0, len(expected))

	if !reflect.DeepEqual(expected, expected) {
		t.Errorf("Expected %v, actual %v", expected, expected)
	}
}

func TestQuickSortEdgeCases(t *testing.T) {
	var arr []int
	var arr2 = []int {1}
	QuickSort(arr, 0, len(arr))
	QuickSort(arr2, 0, len(arr2))

	if !reflect.DeepEqual([]int{}, arr) {
		t.Errorf("Expected %v, actual %v", []int{}, arr)
	}

	if !reflect.DeepEqual([]int{1}, arr2) {
		t.Errorf("Expected %v, actual %v", []int{1}, arr2)
	}
}
