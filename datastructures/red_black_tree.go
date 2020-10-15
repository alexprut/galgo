package datastructures

type RedBlackTree struct {
	root *rbtnode
	size int
}

func NewRedBlackTree() RedBlackTree {
	return RedBlackTree{nil, 0}
}

func (rbt *RedBlackTree) Root() *rbtnode {
	return rbt.root
}

func (rbt *RedBlackTree) Insert(value interface{}) {
	x := &rbtnode{value, true, nil, nil, nil}
	var p *rbtnode
	tmp := rbt.root

	for tmp != nil {
		p = tmp
		if Less(x.value, tmp.value) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	x.parent = p
	if p == nil {
		rbt.root = x
	} else if Less(x.value, p.value) {
		p.left = x
	} else {
		p.right = x
	}
	rbt.insertFixup(x)
	rbt.size++
}

func (rbt *RedBlackTree) insertFixup(x *rbtnode) {
	for x.parent != nil && x.parent.isRed {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y != nil && y.isRed {
				// case 1
				x.parent.isRed = false
				y.isRed = false
				x.parent.parent.isRed = true
				x = x.parent.parent
			} else if x == x.parent.right {
				// case 2
				x = x.parent
				rbt.leftRotation(x)
			} else {
				// case 3
				x.parent.isRed = false
				x.parent.parent.isRed = true
				rbt.rightRotation(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if y != nil && y.isRed {
				// case 1
				x.parent.isRed = false
				y.isRed = false
				x.parent.parent.isRed = true
				x = x.parent.parent
			} else if x == x.parent.left {
				// case 2
				x = x.parent
				rbt.rightRotation(x)
			} else {
				// case 3
				x.parent.isRed = false
				x.parent.parent.isRed = true
				rbt.leftRotation(x.parent.parent)
			}
		}
	}
	rbt.root.isRed = false
}

func (rbt *RedBlackTree) Search(value interface{}) bool {
	return rbt.SearchNode(rbt.root, value) != nil
}

func (rbt *RedBlackTree) SearchNode(root *rbtnode, value interface{}) *rbtnode {
	tmp := root
	for tmp != nil {
		if tmp.value == value {
			return tmp
		}
		if Less(value, tmp.value) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}

	return nil
}

func (rbt *RedBlackTree) Delete(value interface{}) {
	rbt.delete(rbt.SearchNode(rbt.root, value))
	rbt.size--
}

func (rbt *RedBlackTree) delete(z *rbtnode) {
	var x *rbtnode
	y := z
	isOriginalColorRed := y.isRed
	if z.left == nil {
		x = z.right
		rbt.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		rbt.transplant(z, z.left)
	} else {
		y = rbt.MinimumNode(z.right)
		isOriginalColorRed = y.isRed
		x = y.right
		if x != nil && y.parent == z {
			x.parent = y
		} else {
			rbt.transplant(y, y.right)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		}
		rbt.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.isRed = z.isRed
	}
	if x != nil && !isOriginalColorRed {
		rbt.deleteFixup(x)
	}
}

func (rbt *RedBlackTree) deleteFixup(x *rbtnode) {
	for x != rbt.root && !x.isRed {
		if x == x.parent.left {
			w := x.parent.right
			if w.isRed {
				// Case 1
				w.isRed = false
				x.parent.isRed = true
				rbt.leftRotation(x.parent)
				w = x.parent.right
			}
			if w.left != nil && w.right != nil && !w.left.isRed && !w.right.isRed {
				// Case 2
				w.isRed = true
				x = x.parent
			} else if w.right != nil && !w.right.isRed {
				// Case 3
				w.left.isRed = false
				w.isRed = true
				rbt.rightRotation(w)
				w = x.parent.right
			} else {
				// Case 4
				w.isRed = x.parent.isRed
				x.parent.isRed = false
				w.right.isRed = false
				rbt.leftRotation(x.parent)
				x = rbt.root
			}
		} else {
			w := x.parent.left
			if w.isRed {
				// Case 1
				w.isRed = false
				x.parent.isRed = true
				rbt.rightRotation(x.parent)
				w = x.parent.left
			}
			if w.left != nil && w.right != nil && !w.right.isRed && !w.left.isRed {
				// Case 2
				w.isRed = true
				x = x.parent
			} else if w.left != nil && !w.left.isRed {
				// Case 3
				w.right.isRed = false
				w.isRed = true
				rbt.leftRotation(w)
				w = x.parent.left
			} else {
				// Case 4
				w.isRed = x.parent.isRed
				x.parent.isRed = false
				w.left.isRed = false
				rbt.rightRotation(x.parent)
				x = rbt.root
			}
		}
	}
	x.isRed = false
}

func (rbt *RedBlackTree) transplant(u *rbtnode, v *rbtnode) {
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else if u == u.parent.right {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (rbt *RedBlackTree) Minimum() *rbtnode {
	return rbt.MinimumNode(rbt.root)
}

func (rbt *RedBlackTree) Maximum() *rbtnode {
	return rbt.MaximumNode(rbt.root)
}

func (rbt *RedBlackTree) MinimumNode(node *rbtnode) *rbtnode {
	if node == nil || node.left == nil {
		return node
	}
	return rbt.MinimumNode(node.left)
}

func (rbt *RedBlackTree) MaximumNode(node *rbtnode) *rbtnode {
	if node == nil || node.right == nil {
		return node
	}
	return rbt.MaximumNode(node.right)
}

func (rbt *RedBlackTree) Successor(node *rbtnode) *rbtnode {
	if node.right != nil {
		return rbt.MinimumNode(node.right)
	}
	y := node.parent
	for y != nil && node == y.right {
		node = y
		y = y.parent
	}
	return y
}

func (rbt *RedBlackTree) Size() int {
	return rbt.size
}

func (rbt *RedBlackTree) leftRotation(x *rbtnode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		rbt.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (rbt *RedBlackTree) rightRotation(x *rbtnode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		rbt.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

type rbtnode struct {
	value  interface{}
	isRed  bool // when true then color is Red, otherwise Black
	parent *rbtnode
	left   *rbtnode
	right  *rbtnode
}
