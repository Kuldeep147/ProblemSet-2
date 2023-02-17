package main

import "fmt"

func main() {
	fmt.Println("enter the number 6-25 : ")
	var s int
	fmt.Scan(&s)
	for i := 1; i < 11; i++ {
		var t = i * s
		fmt.Printf("%d x %d = %d", s, i, t)
		fmt.Println()
	}
}
