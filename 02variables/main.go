package main

import "fmt"

const Token string = "dlshfdlkf" // public

func main() {
	var username string = "Ritesh"
	fmt.Println(username)
	fmt.Printf("The type of varible is : %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of varible is : %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of varible is : %T\n", smallVal)

	var smallFloat float64 = 43.34335345454
	fmt.Println(smallFloat)
	fmt.Printf("The type of varible is : %T\n", smallFloat)


	// default values and alias
	var anotherVariable int 
	fmt.Println(anotherVariable)

	var anotherString string
	fmt.Println(anotherString)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// walrus operator (can only be used inside methods not in global scope)
	number := 3
	fmt.Println(number)

	fmt.Println(Token)
}
