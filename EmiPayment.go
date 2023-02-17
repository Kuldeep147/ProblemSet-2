package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		p       float64
		rate    float64
		year    float64
		payment float64
	)
	fmt.Println("Enter principal amount : ")
	fmt.Scan(&p)
	fmt.Println("Enter rate of interest : ")
	fmt.Scan(&rate)
	fmt.Println("enter time in years : ")
	fmt.Scan(&year)
	r := rate / 1200
	n := 12 * year
	payment = (p * r) / (1 - 1/math.Pow((1+r), n))
	fmt.Println("Payment is : ", payment)

}
