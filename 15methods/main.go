package main

import "fmt"

func main() {
	ritesh := User{"ritesh", "ritesh@go.dev", true, 16}
	fmt.Println(ritesh)
	fmt.Printf("The values are %+v\n", ritesh)
	fmt.Printf("The Name is %v and Email is %v\n", ritesh.Name, ritesh.Email)

	ritesh.GetStatus()
	ritesh.NewMail()
	fmt.Printf("The Name is %v and Email is %v\n", ritesh.Name, ritesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user online:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is", u.Email)

}