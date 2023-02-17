package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		low       float64 = 6.25
		med       float64 = 7.50
		high      float64 = 9.5
		principal float64 = 1000
	)
	amount := principal * math.Pow((1+low), 2.5)
	amountMed := principal * math.Pow((1+med), 2.5)
	amountHigh := principal * math.Pow((1+high), 2.5)
	fmt.Println("gain low : ", amount)
	fmt.Println("gain medium : ", amountMed)
	fmt.Println("gain high : ", amountHigh)

}
