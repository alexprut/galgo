package datastructures

import "testing"

func TestRBTInsertAndSearch(t *testing.T) {
	tree := NewRedBlackTree()
	for i := 0; i < 100; i++ {
		tree.Insert(i)
		AssertTrue(t, tree.Search(i))
	}
}

func TestRBTInsertAndDeleteAndSearch(t *testing.T) {
	tree := NewRedBlackTree()
	for i := 0; i < 100; i++ {
		tree.Insert(i)
	}

	for i := 0; i < 100; i++ {
		tree.Delete(i)
		AssertFalse(t, tree.Search(i))
	}

	AssertEquals(t, 0, tree.Size())
}

func TestRBTRotateLeft(t *testing.T) {
	tree := NewRedBlackTree()
	x := newrbtreenode(1)
	y := newrbtreenode(2)
	z := newrbtreenode(3)
	w := newrbtreenode(4)
	u := newrbtreenode(5)
	tree.root = x

	x.left = w
	x.right = y
	y.right = z
	y.left = u
	tree.leftRotation(x)

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

func TestRBTRotateRight(t *testing.T) {
	tree := NewRedBlackTree()
	x := newrbtreenode(1)
	y := newrbtreenode(2)
	z := newrbtreenode(3)
	w := newrbtreenode(4)
	u := newrbtreenode(5)
	tree.root = y

	y.left = x
	y.right = z
	x.left = w
	x.right = u
	tree.rightRotation(y)

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

func TestShouldTestInsertFixup(t *testing.T) {
	tree := NewRedBlackTree()
	x := newrbtreenode(11)
	y := newrbtreenode(2)
	y.isRed = true
	z := newrbtreenode(1)
	w := newrbtreenode(7)
	u := newrbtreenode(5)
	u.isRed = true
	a := newrbtreenode(4)
	a.isRed = true
	b := newrbtreenode(8)
	b.isRed = true
	c := newrbtreenode(14)
	d := newrbtreenode(15)
	d.isRed = true
	tree.root = x

	x.left = y
	x.right = c
	y.left = z
	y.right = w
	w.left = u
	w.right = b
	u.left = a
	c.right = d

	tree.insertFixup(a)

	if tree.Root() != nil {
		t.Errorf("Expected %v, actual %v", nil, tree.Root())
	}
	AssertEquals(t, nil, tree.root.parent)
	AssertEquals(t, 15, tree.Maximum().value)
	AssertEquals(t, 1, tree.Minimum().value)
	AssertEquals(t, y, w.left)
	AssertEquals(t, x, w.right)
	AssertEquals(t, z, y.left)
	AssertEquals(t, u, y.right)
	AssertEquals(t, a, u.left)
	AssertEquals(t, b, x.left)
	AssertEquals(t, c, x.right)
	AssertEquals(t, d, c.right)
	AssertEquals(t, nil, z.left)
	AssertEquals(t, nil, z.right)
	AssertEquals(t, nil, a.left)
	AssertEquals(t, nil, a.right)
	AssertEquals(t, nil, b.left)
	AssertEquals(t, nil, b.right)
	AssertEquals(t, nil, d.left)
	AssertEquals(t, nil, d.right)
	AssertEquals(t, nil, u.right)
	AssertEquals(t, nil, c.left)

	AssertEquals(t, false, w.isRed)
	AssertEquals(t, true, y.isRed)
	AssertEquals(t, false, z.isRed)
	AssertEquals(t, false, u.isRed)
	AssertEquals(t, true, a.isRed)
	AssertEquals(t, true, x.isRed)
	AssertEquals(t, false, b.isRed)
	AssertEquals(t, false, c.isRed)
	AssertEquals(t, true, d.isRed)
}

func TestShouldFindSuccessorAndDelete(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(0)

	AssertEquals(t, 1, tree.Successor(tree.Minimum()).value)
	AssertEquals(t, 3, tree.Successor(tree.SearchNode(tree.root, 2)).value)
	AssertEquals(t, 6, tree.Successor(tree.SearchNode(tree.root, 5)).value)
	AssertEquals(t, 8, tree.Successor(tree.SearchNode(tree.root, 7)).value)
	AssertEquals(t, nil, tree.Successor(tree.Maximum()))

	tree.Delete(4)
	tree.Delete(6)
	tree.Delete(1)
	tree.Delete(2)
	tree.Delete(3)
	tree.Delete(5)
	tree.Delete(7)
	tree.Delete(8)
	tree.Delete(9)
	tree.Delete(0)

	AssertEquals(t, 0, tree.Size())
}
