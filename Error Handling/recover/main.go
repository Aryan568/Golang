package main

import "fmt"

func main() {
	fmt.Println("Recover Statement")
	fmt.Println("Start of the program")

	division(2,0)
	division(4,2)

	fmt.Println("End of the program")
}

func division(a, b int) {
	defer func ()  {
		if r:=recover(); r!=nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b==0 {
		panic("Division by zero")
	} else {
		fmt.Println("Result:", a/b)
	}
}

