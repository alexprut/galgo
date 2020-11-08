package datastructures

type Queue struct {
	list DoubleLikedList
}

func NewQueue() Queue {
	return Queue{NewDoubleLinkedList()}
}

func (queue *Queue) Empty() bool {
	return queue.list.Empty()
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.list.InsertFront(value)
}

func (queue *Queue) Dequeue() interface{} {
	return queue.list.RemoveBack()
}

func (queue *Queue) Size() int {
	return queue.list.size
}