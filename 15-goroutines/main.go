package main

import (
	"time" 
	"fmt"
)
func main() { 
	fmt.Println("Concurrency & Goroutines")
	go greeter("Hello World")
	greeter("Hello World")
}

func greeter(msg string) {
	time.Sleep(5*time.Millisecond)
	fmt.Println(msg)
}