package datastructures

import (
	"testing"
)

func TestBasicOperations(t *testing.T) {
	stack := NewStack()

	AssertEquals(t, 0, stack.Size())
	AssertTrue(t, stack.Empty())

	stack.Push(5)
	AssertEquals(t, 1, stack.Size())

	stack.Insert(4)
	stack.Remove()

	stack.Push(4)
	value := stack.Pop()

	AssertEquals(t, 4, value)
	AssertEquals(t, 1, stack.Size())
}
