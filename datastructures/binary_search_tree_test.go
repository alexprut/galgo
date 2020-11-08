package datastructures

import "testing"

func TestInsertAndRetrieveAndSearch(t *testing.T) {
	bst := NewBinarySearchTree()
	AssertEquals(t, 0, bst.Size())

	bst.Insert(1)
	AssertEquals(t, 1, bst.Size())
	AssertTrue(t, bst.Contains(1))
	AssertFalse(t, bst.Contains(2))
	AssertEquals(t, 1, bst.Search(1).value)

	bst.Insert(1)
	AssertEquals(t, 2, bst.Size())
	AssertTrue(t, bst.Contains(1))
	AssertEquals(t, 1, bst.Search(1).value)

	bst.Insert(2)
	AssertEquals(t, 3, bst.Size())
	AssertTrue(t, bst.Contains(2))
	AssertEquals(t, 2, bst.Search(2).value)

	AssertEquals(t, 2, bst.Maximum().value)
	AssertEquals(t, 1, bst.Minimum().value)
}

func TestInsertAndRetrieveAndSearchAndDelete(t *testing.T) {
	bst := NewBinarySearchTree()
	for i := 0; i < 1000; i++ {
		bst.Insert(i)
	}

	for i := 0; i < 1000; i++ {
		AssertEquals(t, i, bst.Search(i).value)
	}

	for i := 0; i < 1000; i++ {
		bst.Delete(bst.Search(i))

		if bst.Search(i) != nil {
			t.Errorf("Expected %v, actual %v", nil, bst.Search(i))
		}
	}

	AssertEquals(t, 0, bst.Size())
}


func TestPreOrderVisit(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(5)

	expected := make([]interface{}, 0)
	expected = append(expected, 4)
	expected = append(expected, 2)
	expected = append(expected, 1)
	expected = append(expected, 3)
	expected = append(expected, 6)
	expected = append(expected, 5)

	AssertEquals(t, expected, bst.PreOrderVisit())
}


func TestInOrderVisit(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(7)

	expected := make([]interface{}, 0)
	expected = append(expected, 1)
	expected = append(expected, 2)
	expected = append(expected, 3)
	expected = append(expected, 4)
	expected = append(expected, 5)
	expected = append(expected, 6)
	expected = append(expected, 7)

	AssertEquals(t, expected, bst.InOrderVisit())
}


func TestPostOrderVisit(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(7)

	expected := make([]interface{}, 0)
	expected = append(expected, 1)
	expected = append(expected, 3)
	expected = append(expected, 2)
	expected = append(expected, 5)
	expected = append(expected, 7)
	expected = append(expected, 6)
	expected = append(expected, 4)

	AssertEquals(t, expected, bst.PostOrderVisit())
}


func TestGetSuccessorAndDelete(t *testing.T){
	bst := NewBinarySearchTree()
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(5)
	bst.Insert(7)

	AssertEquals(t, 5, bst.successor(bst.Search(4)).value)
	AssertEquals(t, 2, bst.successor(bst.Search(1)).value)
	AssertEquals(t, 7, bst.successor(bst.Search(6)).value)
	if bst.successor(bst.Search(7)) != nil {
		t.Errorf("Expected %v, actual %v", nil, bst.successor(bst.Search(7)))
	}

	bst.Delete(bst.Search(6))
	bst.Delete(bst.Search(3))
	bst.Delete(bst.Search(2))
	bst.Delete(bst.Search(4))
	bst.Delete(bst.Search(1))
	bst.Delete(bst.Search(5))
	bst.Delete(bst.Search(7))

	AssertEquals(t, 0, bst.Size())
}