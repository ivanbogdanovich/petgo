package main

import (
	"fmt"

	"petgo/internal/datastruct/cache"
)

type CustomHashMap struct {
	buckets [][]Entry
}

type Entry struct {
	key   string
	value any
}

func NewCustomHashMap(size int) *CustomHashMap {
	return &CustomHashMap{
		buckets: make([][]Entry, size),
	}
}

func hash(key string) int {
	sum := 0

	for _, ch := range key {
		sum += int(ch)
	}

	return sum
}

func (m *CustomHashMap) Set(key string, value any) *CustomHashMap {
	index := hash(key) % len(m.buckets)
	bucket := m.buckets[index]

	entry := Entry{
		key:   key,
		value: value,
	}

	for i, v := range bucket {
		if v.key == entry.key {
			m.buckets[index][i].value = value
			return m
		}
	}

	m.buckets[index] = append(m.buckets[index], entry)
	return m
}

func (m *CustomHashMap) Get(key string) (any, error) {
	index := hash(key) % len(m.buckets)
	bucket := m.buckets[index]

	for _, v := range bucket {
		if v.key == key {
			return v.value, nil
		}
	}

	return nil, fmt.Errorf("key not exist %v", key)
}

func (m *CustomHashMap) Delete(key string) (*CustomHashMap, error) {
	index := hash(key) % len(m.buckets)
	bucket := m.buckets[index]

	for i, v := range bucket {
		if v.key == key {
			m.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return m, nil
		}
	}
	return nil, fmt.Errorf("ket not exist %v", key)
}

func main() {
	stock := map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 8,
	}
	stock["milk"] = 12
	stock["banana"] = 7
	delete(stock, "orange")

	for key, value := range stock {
		fmt.Println("range key", key)
		fmt.Println("range value", value)
	}

	v, ok := stock["orange"]
	fmt.Println("stock", stock)
	fmt.Println("stock", stock["apple"])

	if ok {
		fmt.Println("v", v, ok)
	} else {
		fmt.Println("else", v, ok)
	}

	for key, value := range stock {
		fmt.Println("range key", key)
		fmt.Println("range value", value)
	}

	fmt.Println("len", len(stock))

	var m = map[string]int{}
	m["key"] = 1
	fmt.Println("m", m["key"])

	map1 := NewCustomHashMap(4)

	map1.Set("ananas", 10)
	map1.Set("apple", 10)
	map1.Set("orrange", "lol")
	map1.Set("orrange", "20")
	find, err := map1.Get("orrange")

	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("find %T, value: %v\n", find, find)
	fmt.Println("map1", map1)

	delete, err := map1.Delete("orrange")

	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("delete", delete)
	fmt.Println("map1", map1)

	cache, err := cache.New(2)
	if err != nil {
		fmt.Println("error creating lru cache:", err)
		return
	}

	cache.Set(1, 10)
	cache.Set(2, 20)
	cache.Set(1, 15)
	cache.Set(3, 30)

	if value, ok := cache.Get(1); ok {
		fmt.Println("lru get key=1:", value)
	}

	// if _, ok := cache.Get(2); !ok {
	// 	fmt.Println("lru key=2 evicted")
	// }

	// if value, ok := cache.Get(3); ok {
	// 	fmt.Println("lru get key=3:", value)
	// }

	fmt.Printf("cache items %d\n", cache.Snapshot())
	fmt.Printf("cache list size %d\n", cache.ListSnapshot())
}
