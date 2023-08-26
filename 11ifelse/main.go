package main

import "fmt"

func main() {

	loginCount := 4

	var result string

	if loginCount < 10 {
		result = "Regular Customer"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exact 10"
	}

	fmt.Println(result)


	if 9 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 0; num < 8 {
		fmt.Println("less than count")
	} else {
		fmt.Println("greateer than count")
	}

	// if err != nil {

	// }
}
