package main

import "fmt"

func main() {
	fmt.Println("Choose conversion type..\n 0 for Celsius to Fahrenheit\n 1 for Fahrenheit to celsius..")
	var selector int
	fmt.Scan(&selector)
	var temp float32
	if selector == 0 {
		fmt.Println("Enter temperature in celsius : ")
		fmt.Scan(&temp)
		f := (temp * 9 / 5) + 32
		fmt.Println("Temperature in Fahrenheit is : ", f)
	} else if selector == 1 {
		fmt.Println("Enter temperature in Fahrenheit : ")
		fmt.Scan(&temp)
		c := (temp - 32) * 5 / 9
		fmt.Println("Temperature in celsius is : ", c)
	} else {
		fmt.Println("Enter valid input..")
	}
}
