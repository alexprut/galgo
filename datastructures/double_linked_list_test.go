package datastructures

import "testing"

func TestInsertAndRemoveFront(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	if doubleLinkedList.Head() != nil {
		t.Errorf("Expected %v, actual %v", nil, doubleLinkedList.Head())
	}
	AssertEquals(t, 0, doubleLinkedList.Size())
	doubleLinkedList.InsertFront(5)
	AssertEquals(t, 5, doubleLinkedList.Head().value)
	AssertEquals(t, 1, doubleLinkedList.Size())
	AssertEquals(t, 5, doubleLinkedList.RemoveFront())
	AssertEquals(t, 0, doubleLinkedList.Size())
	if doubleLinkedList.Head() != nil {
		t.Errorf("Expected %v, actual %v", nil, doubleLinkedList.Head())
	}
	doubleLinkedList.InsertBack(1)
	doubleLinkedList.InsertBack(2)
	AssertEquals(t, 2, doubleLinkedList.Size())
	AssertEquals(t, 2, doubleLinkedList.Tail().value)
	AssertEquals(t, 2, doubleLinkedList.RemoveBack())
}

func TestGetNextPrevAndHead(t *testing.T) {
	doubleLinkedList := NewDoubleLinkedList()
	doubleLinkedList.InsertFront(1)
	doubleLinkedList.InsertBack(2)
	doubleLinkedList.InsertBack(3)
	doubleLinkedList.InsertBack(4)
	doubleLinkedList.InsertBack(5)

	AssertEquals(t, 5, doubleLinkedList.RemoveBack())

	AssertEquals(t, 1, doubleLinkedList.RemoveFront())
	
	AssertEquals(t, 3, doubleLinkedList.Size())
	AssertEquals(t, 2, doubleLinkedList.Head().value)
	AssertEquals(t, 3, doubleLinkedList.Head().next.value)
	if doubleLinkedList.Head().prev != nil {
		t.Errorf("Expected %v, actual %v", nil, doubleLinkedList.Head().prev)
	}
}
