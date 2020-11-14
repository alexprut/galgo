package datastructures

import (
	"galgo/algorithms/utils"
	"math"
)

type AVLTree struct {
	root *avlnode
	size int
}

func NewAVLTree() AVLTree {
	return AVLTree{nil, 0}
}

func (avl *AVLTree) Root() *avlnode {
	return avl.root
}

func (avl *AVLTree) Insert(value interface{}) {
	tmp := avl.insert(avl.root, value)
	if avl.root == nil {
		avl.root = tmp
	}
	avl.size++
}

func (avl *AVLTree) insert(node *avlnode, value interface{}) *avlnode {
	if node == nil {
		tmp := newAVLNode(value)
		return &tmp
	}

	if Less(value, node.value) {
		node.left = avl.insert(node.left, value)
		node.left.parent = node
	} else {
		node.right = avl.insert(node.right, value)
		node.right.parent = node
	}

	node.height = utils.MaxInt(avl.height(node.left), avl.height(node.right)) + 1
	balance := avl.getBalance(node)

	// Case: Left Left
	if balance > 1 && Less(value, node.left.value) {
		return avl.rightRotation(node)
	}

	// Case: Right Right
	if balance < -1 && Less(node.right.value, value) {
		return avl.leftRotation(node)
	}

	// Case: Left Right
	if balance > 1 && Less(node.left.value, value) {
		node.left = avl.leftRotation(node.left)
		return avl.rightRotation(node)
	}

	// Case: Right Left
	if balance < -1 && Less(value, node.right.value) {
		node.right = avl.rightRotation(node.right)
		return avl.leftRotation(node)
	}

	return node
}

func (avl *AVLTree) Search(value interface{}) bool {
	return avl.SearchNode(avl.root, value) != nil
}

func (avl *AVLTree) SearchNode(root *avlnode, value interface{}) *avlnode {
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

func (avl *AVLTree) Delete(value interface{}) {
	avl.root = avl.DeleteNode(avl.root, value)
	avl.size--
}

func (avl *AVLTree) DeleteNode(node *avlnode, value interface{}) *avlnode {
	if node == nil {
		return nil
	}

	if Less(value, node.value) {
		node.left = avl.DeleteNode(node.left, value)
	} else if Less(node.value, value) {
		node.right = avl.DeleteNode(node.right, value)
	} else {
		if node.left == nil || node.right == nil {
			if node.left == nil {
				node = node.right
			} else {
				node = node.left
			}
		} else {
			temp := avl.minimum(node.right)
			node.value = temp.value
			node.right = avl.DeleteNode(node.right, temp.value)
		}
	}

	if node == nil {
		return nil
	}

	node.height = int(math.Max(float64(avl.height(node.left)), float64(avl.height(node.right)))) + 1
	balance := avl.getBalance(node)

	// Case: Left Left
	if balance > 1 && avl.getBalance(node.left) >= 0 {
		return avl.rightRotation(node)
	}

	// Case: Left Right
	if balance > 1 && avl.getBalance(node.left) < 0 {
		node.left = avl.leftRotation(node.left)
		return avl.rightRotation(node)
	}

	// Case: Right Right
	if balance < -1 && avl.getBalance(node.right) <= 0 {
		return avl.leftRotation(node)
	}

	// Case: Right Left
	if balance < -1 && avl.getBalance(node.right) > 0 {
		node.right = avl.rightRotation(node.right)
		return avl.leftRotation(node)
	}

	return node
}

func (avl *AVLTree) Minimum() *avlnode {
	return avl.minimum(avl.root)
}

func (avl *AVLTree) Maximum() *avlnode {
	return avl.maximum(avl.root)
}

func (avl *AVLTree) minimum(node *avlnode) *avlnode {
	if node == nil || node.left == nil {
		return node
	}
	return avl.minimum(node.left)
}

func (avl *AVLTree) maximum(node *avlnode) *avlnode {
	if node == nil || node.right == nil {
		return node
	}
	return avl.maximum(node.right)
}

func (avl *AVLTree) Successor(node *avlnode) *avlnode {
	if node.right != nil {
		return avl.minimum(node.right)
	}
	y := node.parent
	for y != nil && node == y.right {
		node = y
		y = y.parent
	}
	return y
}

func (avl *AVLTree) Size() int {
	return avl.size
}

func (avl *AVLTree) leftRotation(x *avlnode) *avlnode {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		avl.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y

	x.height = utils.MaxInt(avl.height(x.left), avl.height(x.right)) + 1
	y.height = utils.MaxInt(avl.height(y.left), avl.height(y.right)) + 1

	return y
}

func (avl *AVLTree) rightRotation(x *avlnode) *avlnode {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		avl.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y

	x.height = int(math.Max(float64(avl.height(x.left)), float64(avl.height(x.right)))) + 1
	y.height = int(math.Max(float64(avl.height(y.left)), float64(avl.height(y.right)))) + 1

	return y
}

func (avl *AVLTree) getBalance(n *avlnode) int {
	if n == nil {
		return 0
	}

	return avl.height(n.left) - avl.height(n.right)
}

func (avl *AVLTree) height(n *avlnode) int {
	if n == nil {
		return 0
	}

	return n.height
}

type avlnode struct {
	value  interface{}
	height int
	parent *avlnode
	left   *avlnode
	right  *avlnode
}

func newAVLNode(value interface{}) avlnode {
	return avlnode{value: value, height: 1, parent: nil, left: nil, right: nil}
}
