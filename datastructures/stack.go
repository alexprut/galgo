package datastructures

type Stack struct {
	list LinkedList
}

func (stack *Stack) Peek() interface{} {
	tmp := stack.Pop()
	stack.Push(tmp)
	return tmp
}

func (stack *Stack) Push(value interface{}) {
	stack.list.InsertFront(value)
}
func (stack *Stack) Size() int {
	return stack.list.Size()
}

func (stack *Stack) Empty() bool {
	return stack.list.Empty()
}

func (stack *Stack) Pop() interface{} {
	return stack.list.RemoveFront()
}

func (stack *Stack) Insert(value interface{}) {
	stack.Push(value)
}

func (stack *Stack) Remove() {
	stack.Pop()
}
