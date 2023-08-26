package main

import "fmt"

func main() {
	greeter()

	addedValue := adder(2 , 3)
	fmt.Println(addedValue)

	proVal, message := proAdder(2, 4, 5, 9)
	fmt.Println(proVal, message)
}

func greeter() {
	fmt.Println("Hello")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value 
	}
	return total, "Result"
}
