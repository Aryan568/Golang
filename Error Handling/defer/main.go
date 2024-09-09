package main

import "fmt"

func main() {
	fmt.Println("Defer statement")

	defer fmt.Println("C")
	fmt.Println("A")
	fmt.Println("B")
}