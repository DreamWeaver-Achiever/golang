package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Wait groups with pointers")
	wg := &sync.WaitGroup{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("First goroutine")
		score = append(score, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Second goroutine")
		score = append(score, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Third goroutine")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Printf("Score: %d\n", score)

}