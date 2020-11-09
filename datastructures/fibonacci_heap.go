package datastructures

import "math"

type FibonacciHeap struct {
	min  *fibonaccinode
	size int
}

func NewFibonacciHeap() FibonacciHeap {
	return FibonacciHeap{nil, 0}
}

func (fbh *FibonacciHeap) Insert(x interface{}) *fibonaccinode {
	node := newfibonaccinode(x)

	if fbh.min == nil {
		fbh.min = &node
	} else {
		if fbh.min.leftSibling != nil {
			node.leftSibling = fbh.min
			node.rightSibling = fbh.min.rightSibling
			fbh.min.rightSibling = &node
			node.rightSibling.leftSibling = &node
		} else {
			fbh.min.leftSibling = &node
		}
		if Less(node.key, fbh.min.key) {
			fbh.min = &node
		}
	}
	fbh.size++
	return &node
}

func (fbh *FibonacciHeap) Minimum() interface{} {
	return fbh.min.key
}

func (fbh *FibonacciHeap) ExtractMin() *fibonaccinode {
	z := fbh.min
	if z != nil {
		if z.child != nil {
			leftChild := z.child.leftSibling
			rightChild := z.child
			z.child.parent = nil
			for leftChild != rightChild {
				leftChild.parent = nil
				leftChild = leftChild.leftSibling
			}
			leftChild = leftChild.rightSibling

			// add child to the root list
			tmp := z.rightSibling
			z.rightSibling = leftChild
			leftChild.leftSibling = z
			tmp.leftSibling = rightChild
			rightChild.rightSibling = tmp
		}

		// remove z from the root list
		z.rightSibling.leftSibling = z.leftSibling
		z.leftSibling.rightSibling = z.rightSibling

		if z == z.rightSibling {
			fbh.min = nil
		} else {
			fbh.min = z.rightSibling
			fbh.consolidate()
		}

		fbh.size--
	}
	return z
}

func (fbh *FibonacciHeap) consolidate() {
	degrees := make([]*fibonaccinode, 0)

	for i := 0; i <= fbh.getDegreeBound(float64(fbh.size)); i++ {
		degrees = append(degrees, nil)
	}
	rootList := make([]*fibonaccinode, 0)
	root := fbh.min
	current := fbh.min.rightSibling
	rootList = append(rootList, root)
	for root != current {
		rootList = append(rootList, current)
		current = current.rightSibling
	}
	for i := 0; i < len(rootList); i++ {
		x := rootList[i]
		d := x.degree
		for degrees[d] != nil {
			y := degrees[d]
			if !Less(x.key, y.key) {
				s := x
				x = y
				y = s
			}
			fbh.link(y, x)
			degrees[d] = nil
			d++
		}
		degrees[d] = x
	}
	fbh.min = nil
	for i := 0; i < len(degrees); i++ {
		degree := degrees[i]
		if degree != nil {
			if fbh.min == nil {
				fbh.min = degree
			} else {
				if Less(degree.key, fbh.min.degree) {
					fbh.min = degree
				}
			}
		}
	}
}

func (fbh *FibonacciHeap) getDegreeBound(n float64) int {
	// The base should be the golden ratio: 1.61803...
	return int(math.Floor(math.Log(n) / math.Log(1.6)))
}

func (fbh *FibonacciHeap) link(y *fibonaccinode, x *fibonaccinode) {
	// Remove y from the root list
	y.leftSibling.rightSibling = y.rightSibling
	y.rightSibling.leftSibling = y.leftSibling
	y.leftSibling = y
	y.rightSibling = y

	// Make y a child of x
	if x.child == nil {
		x.child = y
	} else {
		child := x.child
		y.leftSibling = child
		y.rightSibling = child.rightSibling
		child.rightSibling = y
		y.rightSibling.leftSibling = y
	}
	y.parent = x
	x.degree++
	y.mark = true
}

func (fbh *FibonacciHeap) decreaseKey(x *fibonaccinode, k interface{}) bool {
	if Less(x.key, k) {
		return false
	}
	x.key = k
	y := x.parent
	if y != nil && Less(x.key, y.key) {
		fbh.cut(x, y)
		fbh.cascadingCut(y)
	}
	if Less(x.key, fbh.min.key) {
		fbh.min = x
	}

	return true
}

func (fbh *FibonacciHeap) cut(x *fibonaccinode, y *fibonaccinode) {
	// Remove x from the child list of y
	if x.rightSibling == x {
		y.child = nil
	} else {
		rightSibling := x.rightSibling
		leftSibling := x.leftSibling
		rightSibling.leftSibling = leftSibling
		leftSibling.rightSibling = rightSibling
		y.child = rightSibling
	}
	y.degree--

	// Add x to the root list
	fbh.min.rightSibling.leftSibling = x
	x.rightSibling = fbh.min.rightSibling
	x.leftSibling = fbh.min
	fbh.min.rightSibling = x

	x.parent = nil
	x.mark = false
}

func (fbh *FibonacciHeap) cascadingCut(y *fibonaccinode) {
	z := y.parent
	if z != nil {
		if !y.mark {
			y.mark = true
		} else {
			fbh.cut(y, z)
			fbh.cascadingCut(z)
		}
	}
}

func (fbh *FibonacciHeap) delete(x *fibonaccinode) {
	fbh.decreaseKey(x, fbh.min.key)
	fbh.ExtractMin()
}

type fibonaccinode struct {
	key          interface{}
	degree       int
	mark         bool
	parent       *fibonaccinode
	leftSibling  *fibonaccinode
	rightSibling *fibonaccinode
	child        *fibonaccinode
}

func newfibonaccinode(key interface{}) fibonaccinode {
	node := fibonaccinode{key, 0, false, nil, nil, nil, nil}
	node.leftSibling = &node
	node.rightSibling = &node

	return node
}
