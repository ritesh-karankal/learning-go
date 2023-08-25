package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 34
	
	var ptr = &myNumber

	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

    *ptr = *ptr + 3
	fmt.Println("New Value is", *ptr)
}
