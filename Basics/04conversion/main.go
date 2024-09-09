package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)
<<<<<<< HEAD

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("We have added 1 to your rating, so now the rating is", numRating+1)
	}
=======
	fmt.Printf("Value Type: %T \n", input)

	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("\nWe have added 1 to your rating, so now the rating is", numRating+1)

	


>>>>>>> 8361048 (Golang practice)
}
