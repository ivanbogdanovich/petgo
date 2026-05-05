package queue

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{Value: value}

	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
		q.Size++
		return
	}

	q.Tail.Next = newNode
	q.Tail = newNode
	q.Size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.Head == nil {
		return 0, false
	}

	value := q.Head.Value
	q.Head = q.Head.Next
	q.Size--

	if q.Head == nil {
		q.Tail = nil
	}

	return value, true
}

func (q *Queue) Peek() (int, bool) {
	if q.Head == nil {
		return 0, false
	}
	return q.Head.Value, true
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Len() int {
	return q.Size
}
