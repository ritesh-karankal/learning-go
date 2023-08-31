package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verb - GET")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count: ", byteCount)

	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}
