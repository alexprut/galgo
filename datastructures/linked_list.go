package datastructures

type LikedList struct {
	head *node
	size int
}

func (list *LikedList) InsertFront(value interface{}) {
	list.head = &node{value, list.head}
	list.size++
}

func (list *LikedList) RemoveFront() interface{} {
	list.size--
	tmp := list.head
	list.head = list.head.next
	return tmp.value
}

func (list *LikedList) Head() *node {
	return list.head
}

func (list *LikedList) Empty() bool {
	return list.size == 0
}

func (list *LikedList) Size() int {
	return list.size
}

type node struct {
	value interface{}
	next  *node
}
