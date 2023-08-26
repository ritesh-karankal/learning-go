package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := `You would not believe your eyes
				If ten million fireflies
				Lit up the world as I fell asleep
				'Cause they fill the open air
				And leave teardrops everywhere
				You'd think me rude but I would just stand and stare`

	file, err := os.Create("./lyrics.txt")
	cherNilErr(err)

	length, err := io.WriteString(file, content)
	cherNilErr(err)

	fmt.Println("The length is:", length)

	defer file.Close()

	readFile("./lyrics.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	cherNilErr(err)

	fmt.Println("The data is:\n", string(databyte))
}


func cherNilErr(err error) {
	if err != nil {
		panic(err)
	}
}