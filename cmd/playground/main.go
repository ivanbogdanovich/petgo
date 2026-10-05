package main

import (
	"fmt"
	"os"
	"petgo/internal/datastruct/cache"
	"runtime"
	"sync"
	"time"
)

type Command struct {
	Kind  string
	Reply chan State
}

type State struct {
	Money     int
	Donations int
}

func stateOwner(commands <-chan Command) {
	money := 0
	donations := 0

	for cmd := range commands {
		switch cmd.Kind {
		case "donate":
			money++
			donations++

		case "read":
			cmd.Reply <- State{
				Money:     money,
				Donations: donations,
			}
		}
	}
}

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

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(n)
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

	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}

	wg.Wait()

	money := 0
	donations := 0

	wg1 := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	// go func() {
	// 	for {
	// 		mutex.Lock()
	// 		m := money
	// 		dc := donations
	// 		mutex.Unlock()

	// 		fmt.Println("money read", m, "donations read", dc)

	// 		if m != dc {
	// 			break
	// 		}
	// 	}
	// }()

	wg1.Add(1000)
	for range 1000 {
		go func() {
			defer wg1.Done()

			mutex.Lock()
			money++
			donations++
			mutex.Unlock()
		}()
	}

	fmt.Println("money", money)
	fmt.Println("donations", donations)

	wg1.Wait()
	// version with channel

	// есть переменная money
	// есть канал donateCh
	// 1000 goroutine отправляют сигнал в donateCh
	// одна goroutine-owner читает donateCh и увеличивает money

	newWaitGroup := &sync.WaitGroup{}
	newChannel := make(chan int)

	newWaitGroup.Add(1)
	go func() {
		defer newWaitGroup.Done()
		newChannel <- 100
	}()

	go func() {
		for {
		}
	}()

	go func() {
		for {
		}
	}()

	value, ok := <-newChannel

	close(newChannel)

	newWaitGroup.Wait()

	fmt.Println("value", value, "ok", ok)

	workers := runtime.NumCPU()
	fmt.Println("workers", workers)

	fmt.Println("PID:", os.Getpid())
	fmt.Println("process is alive")
	time.Sleep(60 * 60 * time.Second)
}
