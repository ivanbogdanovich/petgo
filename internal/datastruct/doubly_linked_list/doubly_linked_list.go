package doubly_linked_list

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *List) IsEmpty() bool {
	return l.Size == 0
}

func (l *List) Len() int {
	return l.Size
}

func (l *List) PushFront(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	l.Size++
}

func (l *List) PushBack(value int) {
	newNode := &Node{Value: value}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}

	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
	l.Size++
}

func (l *List) PopFront() (int, bool) {
	if l.Head == nil {
		return 0, false
	}

	value := l.Head.Value

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return value, true
	}

	l.Head = l.Head.Next
	l.Head.Prev = nil
	l.Size--
	return value, true
}

func (l *List) PopBack() (int, bool) {
	if l.Tail == nil {
		return 0, false
	}

	value := l.Tail.Value

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return value, true
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Size--
	return value, true
}
