package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(5)
	if a <= 2 {
		fmt.Println("Head")
	} else if a > 2 {
		fmt.Println("Tail")
	}
}
