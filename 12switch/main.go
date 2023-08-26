package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 or open")
	case 2:
		fmt.Println("Move 2")
	case 3:
		fmt.Println("Move 3")
		fallthrough
	case 4:
		fmt.Println("Move 4")
		fallthrough
	case 5:
		fmt.Println("Move 5")
	case 6:
		fmt.Println("Move 6 and roll again")
	}

}
