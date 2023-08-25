package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning about time package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.January, 1, 2, 53, 3, 2, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
	
}
