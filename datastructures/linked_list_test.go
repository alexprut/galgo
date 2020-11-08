package datastructures

import (
	"testing"
)

func TestShouldSort(t *testing.T) {
	linkedList := LinkedList{}

	if linkedList.Head() != nil {
		t.Errorf("Expected %v, actual %v", nil, linkedList.Head())
	}
	if linkedList.Size() != 0 {
		t.Errorf("Expected %v, actual %v", nil, linkedList.Size())
	}
	linkedList.InsertFront(5)
	if linkedList.Head().value != 5 {
		t.Errorf("Expected %v, actual %v", 5, linkedList.Head().value)
	}
	if linkedList.Size() != 1 {
		t.Errorf("Expected %v, actual %v", 1, linkedList.Size())
	}
	linkedList.RemoveFront()
	if linkedList.Size() != 0 {
		t.Errorf("Expected %v, actual %v", 0, linkedList.Size())
	}
	if linkedList.Head() != nil {
		t.Errorf("Expected %v, actual %v", nil, linkedList.Head())
	}
}
