package datastructures

import "testing"

func TestInsertAndDeleteAndSearch(t *testing.T) {
	tree := NewIntervalTree()
	tree.Insert(0, 3)
	tree.Insert(5, 8)
	tree.Insert(6, 10)
	tree.Insert(8, 9)
	tree.Insert(15, 23)
	tree.Insert(16, 21)
	tree.Insert(17, 19)
	tree.Insert(19, 20)
	tree.Insert(25, 30)
	tree.Insert(26, 26)
	tree.Delete(8, 9)

	if tree.Root() != nil {
		t.Errorf("Expected %v, actual %v", nil, tree.Root())
	}
	AssertTrue(t, tree.Search(0, 3))
	AssertTrue(t, tree.Search(5, 8))
	AssertTrue(t, tree.Search(6, 10))
	AssertFalse(t, tree.Search(8, 9))
	AssertTrue(t, tree.Search(15, 23))
	AssertTrue(t, tree.Search(16, 21))
	AssertTrue(t, tree.Search(17, 19))
	AssertTrue(t, tree.Search(19, 20))
	AssertTrue(t, tree.Search(25, 30))
	AssertTrue(t, tree.Search(26, 26))
}

func TestCheckIfOverlap(t *testing.T) {
	AssertTrue(t, DoOverlap(newIntervalNode(1, 3, false), newIntervalNode(0, 5, false)))
	AssertTrue(t, DoOverlap(newIntervalNode(0, 5, false), newIntervalNode(1, 3, false)))
	AssertTrue(t, DoOverlap(newIntervalNode(1, 3, false), newIntervalNode(0, 2, false)))
	AssertTrue(t, DoOverlap(newIntervalNode(0, 4, false), newIntervalNode(3, 5, false)))
	AssertFalse(t, DoOverlap(newIntervalNode(0, 3, false), newIntervalNode(4, 5, false)))
	AssertFalse(t, DoOverlap(newIntervalNode(4, 5, false), newIntervalNode(0, 3, false)))
}

func TestDoMultipleOperations(t *testing.T) {
	tree := NewIntervalTree()
	tree.Insert(16, 21)
	tree.Insert(8, 9)
	tree.Insert(25, 30)
	tree.Insert(5, 8)
	tree.Insert(15, 23)
	tree.Insert(17, 19)
	tree.Insert(26, 26)
	tree.Insert(0, 3)
	tree.Insert(6, 10)
	tree.Insert(19, 20)

	AssertEquals(t, 1, len(tree.FindAll(0, 2)))
	AssertEquals(t, 0, tree.FindAll(0, 2)[0].low)
	AssertEquals(t, 4, len(tree.FindAll(16, 21)))
	AssertEquals(t, 4, len(tree.FindAll(17, 19)))
	AssertEquals(t, 0, tree.minimum().low)
	AssertEquals(t, 26, tree.maximum().low)
	if tree.Find(4, 9) != nil {
		t.Errorf("Expected %v, actual %v", nil, tree.Find(4, 9))
	}
	if tree.Find(9, 9) != nil {
		t.Errorf("Expected %v, actual %v", nil, tree.Find(9, 9))
	}
	if tree.Find(10, 14) != nil {
		t.Errorf("Expected %v, actual %v", nil, tree.Find(10, 14))
	}
	AssertEquals(t, nil, tree.Find(31, 32))
	AssertEquals(t, nil, tree.Find(-2, -1))
	AssertEquals(t,
		17, tree.successor(tree.SearchNode(tree.Root(), newIntervalNode(16, 21, false))).low)
	AssertEquals(t,
		25, tree.successor(tree.SearchNode(tree.Root(), newIntervalNode(19, 20, false))).low)

	tree.Delete(16, 21)
	tree.Delete(5, 8)
	tree.Delete(17, 19)
	tree.Delete(25, 30)
	tree.Delete(0, 3)
	tree.Delete(8, 9)
	tree.Delete(15, 23)
	tree.Delete(19, 20)
	tree.Delete(26, 26)
	tree.Delete(6, 10)

	AssertEquals(t, 0, tree.Size())
}
