package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter a string : ")
	var word string
	fmt.Scanln(&word)
	if strings.Contains(word, "e") && strings.Contains(word, "a") && strings.Contains(word, "p") {
		fmt.Println("All Present..")
	} else if strings.Contains(word, "e") || strings.Contains(word, "a") || strings.Contains(word, "p") {
		fmt.Println("One or More - Present")
	} else {
		fmt.Println("None Present..")
	}

}
