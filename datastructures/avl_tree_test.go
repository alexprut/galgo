package datastructures

import "testing"

func TestShouldInsertAndSearch(t *testing.T) {
	tree := NewAVLTree()
	for i := 0; i < 100; i++ {
		tree.Insert(i)
		AssertTrue(t, tree.Search(i))
	}

	if tree.root == nil {
		t.Errorf("Expected %v, actual %v", nil, tree.root)
	}
	AssertEquals(t, nil, tree.root.parent)
	AssertEquals(t, 0, tree.Minimum().value)
	AssertEquals(t, 99, tree.Maximum().value)
}

func TestShouldInsertAndDeleteAndSearch(t *testing.T) {
	tree := NewAVLTree()
	for i := 0; i < 100; i++ {
		tree.Insert(i)
	}

	for i := 0; i < 100; i++ {
		tree.Delete(i)
		AssertFalse(t, tree.Search(i))
	}

	AssertEquals(t, 0, tree.size)
}

func TestShouldRotateLeft(t *testing.T) {
	tree := NewAVLTree()
	x := newAVLNode(1)
	y := newAVLNode(2)
	z := newAVLNode(3)
	w := newAVLNode(4)
	u := newAVLNode(5)
	tree.root = &x

	x.left = &w
	x.right = &y
	y.right = &z
	y.left = &u
	tree.leftRotation(&x)

	AssertEquals(t, y, tree.root)
	AssertEquals(t, x, y.left)
	AssertEquals(t, z, y.right)
	AssertEquals(t, w, x.left)
	AssertEquals(t, u, x.right)
	AssertEquals(t, nil, z.left)
	AssertEquals(t, nil, z.right)
	AssertEquals(t, nil, u.left)
	AssertEquals(t, nil, u.right)
	AssertEquals(t, nil, w.left)
	AssertEquals(t, nil, w.right)
}

func TestShouldRotateRight(t *testing.T) {
	tree := NewAVLTree()
	x := newAVLNode(1)
	y := newAVLNode(2)
	z := newAVLNode(3)
	w := newAVLNode(4)
	u := newAVLNode(5)
	tree.root = &y

	y.left = &x
	y.right = &z
	x.left = &w
	x.right = &u
	tree.rightRotation(&y)

	AssertEquals(t, x, tree.root)
	AssertEquals(t, w, x.left)
	AssertEquals(t, y, x.right)
	AssertEquals(t, u, y.left)
	AssertEquals(t, z, y.right)
	AssertEquals(t, nil, w.left)
	AssertEquals(t, nil, w.right)
	AssertEquals(t, nil, u.left)
	AssertEquals(t, nil, u.right)
	AssertEquals(t, nil, z.left)
	AssertEquals(t, nil, z.right)
}

func TestShouldDoMultipleOperations(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(71)
	tree.Insert(34)
	tree.Insert(8)
	tree.Insert(41)
	tree.Insert(68)
	tree.Insert(13)
	tree.Insert(27)

	AssertEquals(t, 7, tree.Size())
	AssertTrue(t, tree.Search(71))
	AssertTrue(t, tree.Search(34))
	AssertTrue(t, tree.Search(8))
	AssertTrue(t, tree.Search(41))
	AssertTrue(t, tree.Search(68))
	AssertTrue(t, tree.Search(13))
	AssertTrue(t, tree.Search(27))

	AssertEquals(t, nil, tree.Successor(tree.SearchNode(tree.root, 71)))
	AssertEquals(t, 41, tree.Successor(tree.SearchNode(tree.root, 34)).value)
	AssertEquals(t, 13, tree.Successor(tree.SearchNode(tree.root, 8)).value)
	AssertEquals(t, 68, tree.Successor(tree.SearchNode(tree.root, 41)).value)
	AssertEquals(t, 71, tree.Successor(tree.SearchNode(tree.root, 68)).value)
	AssertEquals(t, 27, tree.Successor(tree.SearchNode(tree.root, 13)).value)
	AssertEquals(t, 34, tree.Successor(tree.SearchNode(tree.root, 27)).value)

	tree.Delete(34)
	tree.Delete(13)
	tree.Delete(41)
	tree.Delete(71)
	tree.Delete(27)
	tree.Delete(8)
	tree.Delete(68)

	AssertEquals(t, 0, tree.Size())
}
