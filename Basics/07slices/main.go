package main

import "fmt"

func main() {
	fmt.Println("Slices")

	slices:= [5]int{1,2,3,4,5}

	s := slices[2:4]

	fmt.Println(s)
}