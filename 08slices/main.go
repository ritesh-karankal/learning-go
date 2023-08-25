package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of fruitList is: %T", fruitList)

	fruitList = append(fruitList, "Sitafal", "Lichi")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 97
	highScores[1] = 65
	highScores[2] = 23
	highScores[3] = 50
	// highScores[4] = 77

	highScores = append(highScores, 55, 60)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	
	fmt.Println(highScores)

	var courses = []string {"docker", "linux", "reactjs", "networking", "k8s"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
