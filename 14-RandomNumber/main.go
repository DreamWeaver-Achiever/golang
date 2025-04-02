package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Generating Random Numbers Using math/rand and crypto/rand")
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Random number using math/rand: %d\n", rand.Int());

	myRandomNumber, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(987654321))
	if err != nil {
		fmt.Println("Error while generating random number using crypto pkg")
	}
	fmt.Printf("Random number using crypto/rand: %d\n", myRandomNumber)

}