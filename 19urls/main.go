package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hksh4235whj6"

func main() {
	fmt.Println("Handling urls in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)	
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	// fmt.Printf("The query is of type %T\n", qparams)
	// fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	newUrl :=  &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := newUrl.String()
	fmt.Println(anotherUrl)
}
