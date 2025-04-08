package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels And Deadlocks In Golang")

	channel := make(chan string)

	go func() {
		channel <- "Channel Msg"
	}()
	msg := <- channel

	fmt.Printf("Channel message: %s\n", msg)

}