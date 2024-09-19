package main

import "fmt"

// even or odd

func main(){

	arr:= [6]int{0,3,4,5,6,9}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2==0 {
			fmt.Printf("%v:Not Prime\n", arr[i])
		} else {
			fmt.Printf("%v:Prime\n", arr[i])
		}
	}
}