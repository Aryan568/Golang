package main

import "fmt"

func main() {
	fmt.Println("Errorf Function")

	age:= -10

	errMsg:= fmt.Errorf("%v is negative \n Age can't be negative", age)

	if age<0 {
		fmt.Println(errMsg)
	} else {
		fmt.Printf("Age: %v", age)
	}
}