package main

import (
	"fmt"
)

func main() {
	fmt.Println("Rune data type in Golang")

	var mySingleWordRune = "S"
	fmt.Println("Single Word rune: ", mySingleWordRune)

	var myStringRune = "Hello World !!"
	for _,char := range myStringRune {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}