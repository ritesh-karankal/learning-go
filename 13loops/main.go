package main

import "fmt"

func main() {

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("The index is %v, the value is %v\n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto jump
		}

		if rougeValue == 5 {
			// break
			rougeValue++
			continue
			
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

	jump:
		fmt.Println("We jumped here")
}
