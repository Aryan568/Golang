package main

import (
	"fmt"
	"strings"
)

// 6. Implement a function splitString that takes a string and a separator character, returning a slice of strings and the count of substrings.

func splitString(str, separator string) ([]string, int) {
	subStr:= strings.Split(str, separator)

	return subStr, len(subStr)
}

func main() {
	fmt.Println(splitString("Go:Programming", ":"))
	fmt.Println(splitString("Go:Prog:ramming", ":"))
}
