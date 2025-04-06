package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition In Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("First goroutine")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("Second goroutine")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("Third goroutine")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Printf("Score: %d\n", score)

}