package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("The map is", languages)
	fmt.Println("JS", languages["JS"])

	delete(languages, "JS")
	fmt.Println("The map is", languages)

	// loops to iterate maps

	for key, value := range languages {
		fmt.Printf("The key is %v, and the value is %v\n", key, value)
	}
}
