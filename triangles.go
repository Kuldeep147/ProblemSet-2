package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	fmt.Println("Enter side a : ")
	fmt.Scan(&a)
	fmt.Println("Enter side b : ")
	fmt.Scan(&b)
	fmt.Println("Enter side c : ")
	fmt.Scan(&c)
	s1 := a == b
	s2 := b == c
	s3 := c == a
	if s1 && s2 && s3 {
		fmt.Println("Equilateral ")
	} else if !s1 && !s2 && !s3 {
		fmt.Println("Scalene")
	} else {
		fmt.Println("Isosceles")
	}

}
