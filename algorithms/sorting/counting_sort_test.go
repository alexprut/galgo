package sorting

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	var toSort = []int{5, 2, 1, 9, 0, 33, 3, 3, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, CountingSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, CountingSort(toSort))
	}
}

func TestCountingSortInverted(t *testing.T) {
	var toSort = []int{33, 9, 5, 3, 3, 2, 1, 0, 0}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, CountingSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, CountingSort(toSort))
	}
}

func TestCountingSortSorted(t *testing.T) {
	var toSort = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}
	var expected = []int{0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, CountingSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, CountingSort(toSort))
	}
}

func TestCountingSortEdgeCases(t *testing.T) {
	if !reflect.DeepEqual([]int{}, CountingSort([]int{})) {
		t.Errorf("Expected %v, actual %v", []int{}, CountingSort([]int{}))
	}

	if !reflect.DeepEqual([]int{1}, CountingSort([]int{1})) {
		t.Errorf("Expected %v, actual %v", []int{1}, CountingSort([]int{1}))
	}
}
