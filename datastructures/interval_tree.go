package datastructures

import "math"

type IntervalTree struct {
	root *intervalNode
	size int
}

func NewIntervalTree() IntervalTree {
	return IntervalTree{nil, 0}
}

func (it *IntervalTree) Root() *intervalNode {
	return it.root
}

func (it *IntervalTree) Insert(low int, high int) {
	if low > high {
		return
	}
	var x = newIntervalNode(low, high, false)
	var p *intervalNode
	var tmp = it.root

	for tmp != nil {
		p = tmp
		if x.low < tmp.low {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	x.parent = p
	if p == nil {
		it.root = x
	} else if x.low < p.low {
		p.left = x
	} else {
		p.right = x
	}
	it.maxFixup(x)
	it.insertFixup(x)
	it.size++
}

func (it *IntervalTree) insertFixup(x *intervalNode) {
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
				it.leftRotation(x)
			} else {
				// case 3
				x.parent.isRed = false
				x.parent.parent.isRed = true
				it.rightRotation(x.parent.parent)
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
				it.rightRotation(x)
			} else {
				// case 3
				x.parent.isRed = false
				x.parent.parent.isRed = true
				it.leftRotation(x.parent.parent)
			}
		}
	}
	it.root.isRed = false
}

func (it *IntervalTree) Search(low int, high int) bool {
	return it.SearchNode(it.root, newIntervalNode(low, high, false)) != nil
}

func (it *IntervalTree) SearchNode(root *intervalNode, x *intervalNode) *intervalNode {
	tmp := root
	for tmp != nil {
		if tmp.low == x.low && tmp.high == x.high {
			return tmp
		}
		if x.low < tmp.low {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}

	return nil
}

func (it *IntervalTree) Find(low int, high int) *intervalNode {
	return it.FindNode(it.root, newIntervalNode(low, high, false))
}

func (it *IntervalTree) FindNode(root *intervalNode, x *intervalNode) *intervalNode {
	tmp := root
	for tmp != nil && !DoOverlap(tmp, x) {
		if tmp.left != nil && tmp.left.max >= x.low {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}

	return tmp
}

func (it *IntervalTree) FindAll(low int, high int) []*intervalNode {
	return it.FindAllNode(it.root, newIntervalNode(low, high, false))
}

func (it *IntervalTree) FindAllNode(root *intervalNode, x *intervalNode) []*intervalNode {
	var intervals = make([]*intervalNode, 0)
	if root == nil {
		return intervals
	}
	if DoOverlap(root, x) {
		intervals = append(intervals, root)
	}
	if x.low <= root.max {
		var leftIntervals = it.FindAllNode(root.left, x)
		for _, v := range leftIntervals {
			intervals = append(intervals, v)
		}
		var rightIntervals = it.FindAllNode(root.right, x)
		for _, v := range rightIntervals {
			intervals = append(intervals, v)
		}
	}

	return intervals
}

func DoOverlap(a *intervalNode, b *intervalNode) bool {
	return a.low <= b.high && b.low <= a.high
}

func (it *IntervalTree) Delete(low int, high int) {
	it.DeleteNode(it.SearchNode(it.root, newIntervalNode(low, high, false)))
	it.size--
}

func (it *IntervalTree) DeleteNode(z *intervalNode) {
	var x *intervalNode
	y := z
	isOriginalColorRed := y.isRed
	if z.left == nil {
		x = z.right
		it.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		it.transplant(z, z.left)
	} else {
		y = it.minimumNode(z.right)
		isOriginalColorRed = y.isRed
		x = y.right
		if x != nil && y.parent == z {
			x.parent = y
		} else {
			it.transplant(y, y.right)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		}
		it.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.isRed = z.isRed
	}
	it.maxFixup(y)
	if x != nil && !isOriginalColorRed {
		it.deleteFixup(x)
	}
}

func (it *IntervalTree) deleteFixup(x *intervalNode) {
	for x != it.root && !x.isRed {
		if x == x.parent.left {
			w := x.parent.right
			if w.isRed {
				// Case 1
				w.isRed = false
				x.parent.isRed = true
				it.leftRotation(x.parent)
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
				it.rightRotation(w)
				w = x.parent.right
			} else {
				// Case 4
				w.isRed = x.parent.isRed
				x.parent.isRed = false
				w.right.isRed = false
				it.leftRotation(x.parent)
				x = it.root
			}
		} else {
			w := x.parent.left
			if w.isRed {
				// Case 1
				w.isRed = false
				x.parent.isRed = true
				it.rightRotation(x.parent)
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
				it.leftRotation(w)
				w = x.parent.left
			} else {
				// Case 4
				w.isRed = x.parent.isRed
				x.parent.isRed = false
				w.left.isRed = false
				it.rightRotation(x.parent)
				x = it.root
			}
		}
	}
	x.isRed = false
}

func (it *IntervalTree) transplant(u *intervalNode, v *intervalNode) {
	if u.parent == nil {
		it.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else if u == u.parent.right {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (it *IntervalTree) minimum() *intervalNode {
	return it.minimumNode(it.root)
}

func (it *IntervalTree) maximum() *intervalNode {
	return it.maximumNode(it.root)
}

func (it *IntervalTree) minimumNode(node *intervalNode) *intervalNode {
	if node == nil || node.left == nil {
		return node
	}
	return it.minimumNode(node.left)
}

func (it *IntervalTree) maximumNode(node *intervalNode) *intervalNode {
	if node == nil || node.right == nil {
		return node
	}
	return it.maximumNode(node.right)
}

func (it *IntervalTree) successor(node *intervalNode) *intervalNode {
	if node.right != nil {
		return it.minimumNode(node.right)
	}
	y := node.parent
	for y != nil && node == y.right {
		node = y
		y = y.parent
	}
	return y
}

func (it *IntervalTree) Size() int {
	return it.size
}

func (it *IntervalTree) maxCalculate(x *intervalNode) {
	x.max = x.high
	if x != nil && x.left != nil {
		x.max = int(math.Max(float64(x.max), float64(x.left.max)))
	}
	if x != nil && x.right != nil {
		x.max = int(math.Max(float64(x.max), float64(x.right.max)))
	}
}

func (it *IntervalTree) maxFixup(x *intervalNode) {
	if x != nil {
		it.maxCalculate(x)
		it.maxFixup(x.parent)
	}
}

func (it *IntervalTree) leftRotation(x *intervalNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		it.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
	it.maxCalculate(x)
	it.maxCalculate(x.parent)
}

func (it *IntervalTree) rightRotation(x *intervalNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		it.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
	it.maxCalculate(x)
	it.maxCalculate(x.parent)
}

type intervalNode struct {
	low    int
	high   int
	max    int
	isRed  bool
	parent *intervalNode
	left   *intervalNode
	right  *intervalNode
}

func newIntervalNode(low int, high int, isRed bool) *intervalNode {
	return &intervalNode{
		low:    low,
		high:   high,
		max:    high,
		isRed:  isRed,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}
