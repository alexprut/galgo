package datastructures

type Comparator interface {
	Compare(b interface{}) int
}

func Less(a, b interface{}) bool {
	switch a.(type) {
	case int:
		if ai, ok := a.(int); ok {
			if bi, ok := b.(int); ok {
				return ai < bi
			}
		}
	case string:
		if ai, ok := a.(string); ok {
			if bi, ok := b.(string); ok {
				return ai < bi
			}
		}
	case Comparator:
		if ai, ok := a.(Comparator); ok {
			if bi, ok := b.(Comparator); ok {
				return ai.Compare(bi) < 0
			}
		}
	default:
		panic("Unknown")
	}
	return false
}

type BinarySearchTree struct {
	root *binarynode
	size int
}

func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{nil, 0}
}

func (bst *BinarySearchTree) Insert(value interface{}) {
	node := binarynode{}
	node.value = value

	if bst.root == nil {
		bst.root = &node
	} else {
		bst.Insert(node)
	}
	bst.size++
}

func (bst *BinarySearchTree) Search(value interface{}) *binarynode {
	return bst.search(bst.root, value)
}

func (bst *BinarySearchTree) Delete(node *binarynode) {
	if node.left == nil {
		bst.transplant(node, node.right)
	} else if node.right == nil {
		bst.transplant(node, node.left)
	} else {
		y := bst.minimum(node.right)
		if y.parent != node {
			bst.transplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		}
		bst.transplant(node, y)
		y.left = node.left
		y.left.parent = y
	}
	bst.size--
}

func (bst *BinarySearchTree) transplant(u *binarynode, v *binarynode) {
	if u.parent == nil {
		bst.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (bst *BinarySearchTree) successor(x *binarynode) *binarynode {
	if x.right != nil {
		return bst.minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

func (bst *BinarySearchTree) PreOrderVisit() []interface{} {
	if bst.root == nil {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, 0)
	stack := NewStack()
	stack.Push(bst.root)
	for !stack.Empty() {
		current := stack.Pop().(*binarynode)
		result = append(result, current.value)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}

	return result
}

func (bst *BinarySearchTree) InOrderVisit() []interface{} {
	if bst.root == nil {
		return make([]interface{}, 0)
	}
	node := bst.root
	result := make([]interface{}, 0)
	stack := NewStack()
	for !stack.Empty() || node != nil {
		if node != nil {
			stack.Push(node)
			node = node.left
		} else {
			node = stack.Pop().(*binarynode)
			result = append(result, node.value)
			node = node.right
		}
	}

	return result
}

func (bst *BinarySearchTree) PostOrderVisit() []interface{} {
	if bst.root == nil {
		return make([]interface{}, 0)
	}
	result := make([]interface{}, 0)
	stack := NewStack()
	node := bst.root
	var lastVisited *binarynode
	for !stack.Empty() || node != nil {
		if node != nil {
			stack.Push(node)
			node = node.left
		} else {
			tmp := stack.Peek().(*binarynode)
			if tmp.right != nil && lastVisited != tmp.right {
				node = tmp.right
			} else {
				result = append(result, tmp.value)
				lastVisited = stack.Pop().(*binarynode)
			}
		}
	}

	return result
}

func (bst *BinarySearchTree) search(node *binarynode, value interface{}) *binarynode {
	if node == nil || node.value == value {
		return node
	}
	if Less(node.value, value) {
		return bst.search(node.right, value)
	}
	return bst.search(node.left, value)
}

func (bst *BinarySearchTree) Contains(value interface{}) bool {
	return bst.search(bst.root, value) != nil
}

func (bst *BinarySearchTree) Size() int {
	return bst.size
}

func (bst *BinarySearchTree) insert(node *binarynode) {
	var y *binarynode
	x := bst.root
	for x != nil {
		y = x
		if Less(node.value, x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = y
	if y == nil {
		bst.root = node
	} else if Less(node.value, y.value) {
		y.left = node
	} else {
		y.right = node
	}
}

func (bst *BinarySearchTree) Minimum() *binarynode {
	return bst.minimum(bst.root)
}

func (bst *BinarySearchTree) Maximum() *binarynode {
	return bst.maximum(bst.root)
}

func (bst *BinarySearchTree) minimum(node *binarynode) *binarynode {
	if node == nil || node.left == nil {
		return node
	}
	return bst.minimum(node.left)
}

func (bst *BinarySearchTree) maximum(node *binarynode) *binarynode {
	if node == nil || node.right == nil {
		return node
	}
	return bst.maximum(node.right)
}

type binarynode struct {
	value  interface{}
	parent *binarynode
	left   *binarynode
	right  *binarynode
}
