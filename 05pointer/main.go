package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	num := 20
	var ptr = &num

	fmt.Println("Pointer address of num is", ptr)
	fmt.Println("Value of num by pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Updated num value is", num)
}
