package main

import (
	"petgo/internal/datastruct/doubly_linked_list"
	"petgo/internal/datastruct/queue"
	"petgo/internal/datastruct/stack"
)

func main() {
	doubly_linked_list.CallDoublyLinkedList()
	queue.CallQueue()
	stack.CallStack()
}