package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("New Function")
	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Go" {
		fmt.Println(myError)
	}
}
