package main

import "fmt"

// 2. Create a function divideAndRemainder that takes two integers as parameters and returns their quotient and remainder.

func divAndRem(a, b int) (int, int) {
	// var q float32
	// q= float32(a)/float32(b)
	q:=a/b
	r:=a%b

	return q, r
}

func main() {
	fmt.Println(divAndRem(12,7))
}