package main

import "fmt"

func main() {
	ritesh := User{"ritesh", "ritesh@go.dev", true, 16}
	fmt.Println(ritesh)
	fmt.Printf("The values are %+v\n", ritesh)
	fmt.Printf("The Name is %v and Email is %v", ritesh.Name, ritesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
