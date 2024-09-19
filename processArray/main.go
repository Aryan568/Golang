package main

import "fmt"

// 8. Create a function processArray that takes a slice of integers and returns the sum, average, and a slice containing only even numbers.

func processArray(a []int) (int, float32, []int) {
	sum:=0
	res:=[]int{}
	for i := 0; i < len(a); i++ {
		sum+= a[i]
		if a[i]%2==0 {
			res=append(res, a[i])
		}
	} 
	
	avg:= float32(sum)/float32(len(a))

	return sum, avg, res
}

func main() {
	slice:= []int{2,4,5,8,1}
	fmt.Println(processArray(slice))
}