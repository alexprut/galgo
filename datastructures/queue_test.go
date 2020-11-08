package datastructures

import (
	"testing"
)

func TestBasicQueueOperations(t *testing.T) {
	queue := NewQueue()
	AssertEquals(t, 0, queue.Size())
	AssertTrue(t, queue.Empty())

	queue.Enqueue(5)
	AssertEquals(t, 1, queue.Size())

	queue.Enqueue(4)
	value := queue.Dequeue()

	AssertEquals(t, 5, value)
	AssertEquals(t, 1, queue.Size())
}