package main

import (
	rands "crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	ints, _ := rands.Int(rands.Reader, big.NewInt(2))
	intst := fmt.Sprintf("%d", ints)
	log.Println(intst)
}
