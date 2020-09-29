package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var toSort = []int{5, 2, 1, 9, 0, 33, 3, 3, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, InsertionSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, InsertionSort(toSort))
	}
}

func TestInsertionSortInverseSorted(t *testing.T) {
	var toSort = []int{33, 9, 5, 3, 3, 2, 1, 0, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, InsertionSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, InsertionSort(toSort))
	}
}

func TestInsertionSortSorted(t *testing.T) {
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, InsertionSort(expected)) {
		t.Errorf("Expected %v, actual %v", expected, InsertionSort(expected))
	}
}

func TestInsertionSortEdgeCases(t *testing.T) {
	if !reflect.DeepEqual([]int{}, InsertionSort([]int{})) {
		t.Errorf("Expected %v, actual %v", []int{}, InsertionSort([]int{}))
	}

	if !reflect.DeepEqual([]int{1}, InsertionSort([]int{1})) {
		t.Errorf("Expected %v, actual %v", []int{1}, InsertionSort([]int{1}))
	}
}