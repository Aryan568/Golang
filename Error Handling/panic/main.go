package main

import "fmt"

func main() {
	fmt.Println("Panic Statement")

	panic("There is an error!!")
	fmt.Println("Waiting for execution")
}
