package main

import "fmt"

func main() {
	fmt.Println("enter the nth number : ")
	var n int
	fmt.Scan(&n)
	fmt.Printf("H%d = ", n)
	for i := 1; i <= n; i++ {
		fmt.Printf("1/%d", i)
		if i < n {
			fmt.Printf(" + ")
		}
	}
	fmt.Println()
}
