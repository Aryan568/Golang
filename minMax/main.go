package main

import (
	"fmt"
)

// 4. Write a function minMax that takes a variadic parameter of integers and returns the minimum and maximum values from the input.

func minMax(input ...int) (int, int) {
	min:= input[0]
	max:=input[0]
	for i := 1; i < len(input)-1; i++ {
		if min>input[i+1] {
			min=input[i]
		}
		if max<input[i+1] {
			max=input[i]
		}
	}
	return min, max
}

func main() {
	fmt.Println(minMax(2, -9, 0, 0, 5, 5, 44, 77))
}
