package main

import "fmt"

// 7. Write a function fibonacci that takes an integer n and returns the nth Fibonacci number and a boolean indicating whether the calculation was successful.

func fib(n int) (int, bool) {
	if n<0 {
		return 0, false
	}
	if n <= 1 {
		return n, true
	}
	var n1, n2 = 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2, true
}

func main() {
	fmt.Println(fib(9))  
	fmt.Println(fib(0))
	fmt.Println(fib(-9))    
}