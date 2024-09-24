package main

import (
	"fmt"
	"strings"
)

// 6. Implement a function splitString that takes a string and a separator character, returning a slice of strings and the count of substrings.
// 9. Implement a function parseCoordinates that takes a string in the format "x,y" and returns two floats (x and y) and an error if the format is invalid.
// 10. Write a function generatePassword that takes desired length and whether to include special characters as parameters, returning a generated password and its strength rating.
// 11. Write a func that should check password strength(len>8&&hasSpecialChar&&hasUppernLowerCase&&NumericNum(Very Strong))

func splitString(str, separator string) ([]string, int) {
	subStr:= strings.Split(str, separator)

	return subStr, len(subStr)
}

func main() {
	fmt.Println(splitString("Go:Programming", ":"))
	fmt.Println(splitString("Go:Prog:ramming", ":"))
}
