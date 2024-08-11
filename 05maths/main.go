package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Math in Go!")

	// fmt.Println(rand.Intn(7))

	cryptoRandom, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(cryptoRandom)

}
