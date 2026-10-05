package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Semaphore interface {
	Acquire(context.Context, int64) error
	TryAcquire(int64) bool
	Release(int64)
}

type semaphore struct {
	ch chan struct{}
}

func NewSemaphore(size int64) Semaphore {
	if size <= 0 {
		panic("semaphore size must be greater than 0")
	}

	return &semaphore{
		ch: make(chan struct{}, size),
	}
}

func (s *semaphore) Acquire(ctx context.Context, n int64) error {
	if n <= 0 {
		return fmt.Errorf("acquire value must be greater than 0")
	}

	for i := int64(0); i < n; i++ {
		select {
		case s.ch <- struct{}{}:

		case <-ctx.Done():
			for j := int64(0); j < i; j++ {
				<-s.ch
			}

			return ctx.Err()
		}
	}

	return nil
}

func (s *semaphore) TryAcquire(n int64) bool {
	if n <= 0 {
		return false
	}

	acquired := int64(0)

	for acquired < n {
		select {
		case s.ch <- struct{}{}:
			acquired++

		default:
			// returning already acquired
			for i := int64(0); i < acquired; i++ {
				<-s.ch
			}

			return false
		}
	}

	return true
}

func (s *semaphore) Release(n int64) {
	if n <= 0 {
		panic("release value must be greater than 0")
	}

	for i := int64(0); i < n; i++ {
		select {
		case <-s.ch:
			// slot released

		default:
			panic("semaphore: released more than acquired")
		}
	}
}

func main() {
	sem := NewSemaphore(2)

	ctx := context.Background()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Println("acquire error:", err)
				return
			}

			defer sem.Release(1)

			fmt.Printf("goroutine %d start\n", id)

			time.Sleep(2 * time.Second)

			fmt.Printf("goroutine %d finish\n", id)
		}(i)
	}

	wg.Wait()
}
