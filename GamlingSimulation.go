package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//stake, goal and trails
	fmt.Println("Enter stake : ")
	var (
		stake   int
		goal    int
		trails  int
		matches float32
		won     float32
		avg     float32
	)
	fmt.Scan(&stake)
	fmt.Println("enter goal : ")
	fmt.Scan(&goal)
	fmt.Println("Enter trails : ")
	fmt.Scan(&trails)
	for stake < goal && stake > 0 {
		bet := rand.Float32()
		stake -= trails
		matches++
		if bet < 0.5 {
			stake += 2 * trails
			won++
		}
	}
	avg = won / matches
	fmt.Println("Average matches won : ", avg)
	fmt.Println("Final stake : ", stake)

}
