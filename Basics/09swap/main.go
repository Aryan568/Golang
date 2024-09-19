package main

import "fmt"

func main() {
	fmt.Println("Swap")

	fmt.Printf("Enter number to swap\n")
	var a,b,temp int
	fmt.Scanf("%d \t %d", &a, &b)

	fmt.Printf("Numbers before swapping: a=%d and b=%d\n", a,b)

	temp=a
	a=b
	b=temp

	fmt.Printf("Number after swapping: a=%d and b=%d",a,b)

}