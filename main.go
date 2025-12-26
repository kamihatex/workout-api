package main

import (
	"fmt"
)

func main() {
	fmt.Println("first go program ever :D!")

	// variables

	var name string = "Kami"
	fmt.Printf("this is my name %s\n", name)

	age := 26
	fmt.Printf("and this is my age %d\n", age)

	// loops

	for i := 0; i < 5; i++ {
		fmt.Println("this is i: ", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("this is counter: ", counter)
		counter++
	}

	// arrays and slices

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("these are our numbers %v\n", numbers)
}
