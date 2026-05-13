package cache

import (
	"fmt"

	dll "petgo/internal/datastruct/doubly_linked_list"
)

type entry struct {
	value int
	node  *dll.Node
}

type Cache struct {
	Capacity int
	List     dll.List
	Items    map[int]*entry
}

func New(capacity int) (*Cache, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("capacity must be greater than zero")
	}

	return &Cache{
		Capacity: capacity,
		Items:    make(map[int]*entry, capacity),
	}, nil
}

func (c *Cache) Get(key int) (int, bool) {
	e, ok := c.Items[key]
	if !ok {
		return 0, false
	}

	c.moveToFront(e.node)
	return e.value, true
}

func (c *Cache) Set(key, value int) {
	if e, ok := c.Items[key]; ok {
		e.value = value
		c.moveToFront(e.node)
		return
	}

	node := &dll.Node{Value: key}
	c.pushFront(node)
	c.Items[key] = &entry{
		value: value,
		node:  node,
	}

	if len(c.Items) > c.Capacity {
		c.evictLeastRecentlyUsed()
	}
}

func (c *Cache) Clear() {
	for key := range c.Items {
		delete(c.Items, key)
	}
	c.List = dll.List{}
}

func (c *Cache) Snapshot() map[int]int {
	out := make(map[int]int, len(c.Items))
	for k, e := range c.Items {
		if e != nil {
			out[k] = e.value
		}
	}
	return out
}

func (c *Cache) ListSnapshot() []int {
	values := make([]int, 0, c.List.Size)
	for cur := c.List.Head; cur != nil; cur = cur.Next {
		values = append(values, cur.Value)
	}
	return values
}

func (c *Cache) moveToFront(node *dll.Node) {
	if node == nil || c.List.Head == node {
		return
	}

	c.removeNode(node)
	c.pushFront(node)
}

func (c *Cache) pushFront(node *dll.Node) {
	node.Prev = nil
	node.Next = c.List.Head

	if c.List.Head != nil {
		c.List.Head.Prev = node
	} else {
		c.List.Tail = node
	}

	c.List.Head = node
	c.List.Size++
}

func (c *Cache) removeNode(node *dll.Node) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		c.List.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		c.List.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
	c.List.Size--
}

func (c *Cache) evictLeastRecentlyUsed() {
	lru := c.List.Tail
	if lru == nil {
		return
	}

	c.removeNode(lru)
	delete(c.Items, lru.Value)
}
