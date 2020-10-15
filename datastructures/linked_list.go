package datastructures

type LinkedList struct {
	head *listnode
	size int
}
func NewLinkedList() LinkedList {
	return LinkedList{nil, 0}
}

func (list *LinkedList) InsertFront(value interface{}) {
	list.head = &listnode{value, list.head}
	list.size++
}

func (list *LinkedList) RemoveFront() interface{} {
	list.size--
	tmp := list.head
	list.head = list.head.next
	return tmp.value
}

func (list *LinkedList) Head() *listnode {
	return list.head
}

func (list *LinkedList) Empty() bool {
	return list.size == 0
}

func (list *LinkedList) Size() int {
	return list.size
}

type listnode struct {
	value interface{}
	next  *listnode
}
