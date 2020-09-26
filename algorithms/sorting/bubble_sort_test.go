package sorting

import (
	"reflect"
	"testing"
)

func TestShouldSort(t *testing.T) {
	var toSort = []int {5, 2, 1, 9, 0, 33, 3, 3, 0}
	var expected = []int {0, 0, 1, 2, 3, 3, 5, 9, 33}

	if !reflect.DeepEqual(expected, BubbleSort(toSort)) {
		t.Errorf("Expected %v, actual %v", expected, BubbleSort(toSort))
	}
}
