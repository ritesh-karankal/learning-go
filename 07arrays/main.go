package main

import "fmt"

func main() {
	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[4] = "Mango"

	fmt.Println(fruitList)

	var rollNo = [5]int {3, 5, 3, 2}
	fmt.Println(rollNo)

	fmt.Println("len of rollNo", len(rollNo))
}
