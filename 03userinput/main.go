package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hellos := "Hello there!"
	fmt.Println(hellos)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	input, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you ", input)
}
