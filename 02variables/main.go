package main

import "fmt"

const value uint8 = 1

var name string = "James"

func main() {

	var a, b uint8 = 5, 6
	var username string = "Aryan"
	fmt.Println(username)
	fmt.Println("My name is ", username)
	fmt.Printf("Variable is of %v type\n", username)
	fmt.Printf("Variable is of %T type\n", username)
	fmt.Println(name)
	fmt.Println(value)
	fmt.Println(a, b)
}
