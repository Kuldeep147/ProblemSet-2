package main

import "fmt"

func main() {
	fmt.Println("Enter the integer number : ")
	var s int
	fmt.Scan(&s)
	var t bool
	if s%3 == 0 {
		fmt.Printf("Pling")
		t = true
	}
	if s%5 == 0 {
		fmt.Printf("Plang")
		t = true
	}
	if s%7 == 0 {
		fmt.Printf("Plong")
		t = true
	}
	if t == false {
		fmt.Println(s)
	}
	fmt.Println()
}
