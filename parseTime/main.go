package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 3. Implement a function parseTime that accepts a string in the format "HH:MM" and returns three integers: hours, minutes, and an error (if the input is invalid).

func parseTime(str string) (int, int, error) {
	parts:= strings.Split(str, ":")

	if len(parts)!=2 {
		return 0, 0, errors.New("wrong format")
	}

	hrs, err:= strconv.Atoi(parts[0])
	if err!=nil|| hrs<0|| hrs>23 {
		return 0,0, errors.New("invalid hrs")
	}

	min, err:= strconv.Atoi(parts[1])
	if err!=nil|| min<0|| min>59 {
		return 0,0, errors.New("invalid min")
	}
	return hrs, min, nil
}

func main() {
	fmt.Println(parseTime("11:45"))
}