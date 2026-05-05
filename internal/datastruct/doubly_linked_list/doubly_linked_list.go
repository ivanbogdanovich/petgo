package doubly_linked_list

import "fmt"

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

func (l *List) PrintForward() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Value)
	}
	fmt.Println()
}

func (l *List) PrintBackward() {
	for cur := l.Tail; cur != nil; cur = cur.Prev {
		fmt.Printf("%d ", cur.Value)
	}
	fmt.Println()
}

func CallDoublyLinkedList() {
	var list List

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	fmt.Println("forward:")
	list.PrintForward()

	fmt.Println("backward:")
	list.PrintBackward()

	list.PushFront(5)
	list.PushFront(1)

	fmt.Println("after PushFront:")
	list.PrintForward()
	list.PrintBackward()

	v, ok := list.PopFront()
	fmt.Println("PopFront:", v, ok)
	list.PrintForward()

	v, ok = list.PopBack()
	fmt.Println("PopBack:", v, ok)
	list.PrintForward()

	fmt.Println("len:", list.Len())
	fmt.Println("empty:", list.IsEmpty())
}
