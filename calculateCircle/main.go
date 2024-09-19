package main

import "fmt"

// 1. Write a function called calculateCircle that takes the radius as a parameter and returns both the area and circumference of the circle.

func calculateCircle(r float32) (float32, float32) {
	const pi float32 = 3.14

	circum := 2 * pi * r
	area := pi * r * r

	return circum, area
}

func main() {
	fmt.Println(calculateCircle(4))
}
