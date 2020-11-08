package utils

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	var arr = []int{1, 2, 0, 9, 5}
	var expected = []int{0, 2, 1, 9, 5}

	Swap(arr, 0, 2)

	if !reflect.DeepEqual(expected, arr) {
		t.Errorf("Expected %v, actual %v", expected, arr)
	}
}

func TestMax(t *testing.T) {
	var arr = []int{1, 2, 0, 9, 5}

	if 9 != Max(arr) {
		t.Errorf("Expected %v, actual %v", 9, Max(arr))
	}
}