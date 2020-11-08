package datastructures

type DoubleLikedList struct {
	head *node
	tail *node
	size int
}

func NewDoubleLinkedList() DoubleLikedList {
	return DoubleLikedList{nil, nil, 0}
}

func (list *DoubleLikedList) InsertFront(value interface{}) {
	node := &node{value, list.Head(), nil}
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}

	list.size++
}

func (list *DoubleLikedList) InsertBack(value interface{}) {
	node := &node{value, nil, list.tail}
	if list.head == nil {
		list.head = node
	} else {
		node.prev = list.tail
		list.tail.next = node
	}
	list.tail = node

	list.size++
}

func (list *DoubleLikedList) RemoveFront() interface{} {
	list.size--

	if list.head == list.tail {
		node := list.head
		list.head = nil
		list.tail = nil
		return node.value
	}

	node := list.head
	list.head = node.next
	list.head.prev = nil
	return node.value
}

func (list *DoubleLikedList) RemoveBack() interface{} {
	if list.head == list.tail {
		node := list.head
		list.head = nil
		list.tail = nil
		return node.value
	}

	node := list.tail
	node.prev.next = nil
	list.tail = node.prev
	list.size--
	return node.value
}

func (list *DoubleLikedList) Head() *node {
	return list.head
}

func (list *DoubleLikedList) Tail() *node {
	return list.tail
}

func (list *DoubleLikedList) Empty() bool {
	return list.size == 0
}

func (list *DoubleLikedList) Size() int {
	return list.size
}

type node struct {
	value interface{}
	next  *node
	prev  *node
}
