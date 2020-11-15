package datastructures

import (
	"math/rand"
	"testing"
)

func TestShouldInsertAndGetMinimum(t *testing.T) {
	heap := NewFibonacciHeap()
	heap.Insert(1)
	AssertEquals(t, 1, heap.size)
	AssertEquals(t, 1, heap.Minimum())
}

func TestShouldInsertAndExtractMinimum(t *testing.T) {
	heap := NewFibonacciHeap()
	heap.Insert(1)
	AssertEquals(t, 1, heap.ExtractMin().key)
	AssertEquals(t, 0, heap.size)
}

func TestShouldDecreaseKey(t *testing.T) {
	heap := NewFibonacciHeap()
	x := heap.Insert(2)
	heap.decreaseKey(x, 1)
	AssertEquals(t, 1, x.key)
}

func TestShouldDeleteNode(t *testing.T) {
	heap := NewFibonacciHeap()
	x := heap.Insert(2)
	heap.delete(x)
	AssertEquals(t, 0, heap.size)
	AssertEquals(t, nil, heap.min)
}

func TestFibonacciHeapMultipleOperations(t *testing.T) {
	nodes := make([]*fibonaccinode, 0)
	heap := NewFibonacciHeap()
	for i := 0; i < 6; i++ {
		nodes = append(nodes, heap.Insert(i))
	}
	for i := 0; i < len(nodes); i++ {
		heap.decreaseKey(nodes[i], nodes[i].key.(int)-1)
	}
	AssertEquals(t, -1, heap.Minimum())

	heap.delete(nodes[0])
	heap.delete(nodes[1])
	heap.delete(nodes[4])
	heap.delete(nodes[5])

	AssertEquals(t, 1, heap.Minimum())
	AssertEquals(t, 2, heap.size)
}

func TestShouldDoMultipleRandomOperations(t *testing.T) {
	nodes := make([]*fibonaccinode, 0)
	heap := NewFibonacciHeap()

	for i := 0; i < 100; i++ {
		nodes = append(nodes, heap.Insert(i*rand.Intn(1)))
	}
	for i := 0; i < len(nodes)*rand.Intn(1); i++ {
		heap.decreaseKey(nodes[i], nodes[i].key.(int)-1)
	}
	for i := 0; i < len(nodes)*rand.Intn(1); i++ {
		heap.delete(nodes[i])
	}
}
