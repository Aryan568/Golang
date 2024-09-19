package main

import (
	"errors"
	"fmt"
)

// 5. Create a function createPerson that takes name and age as parameters and returns a struct Person and an error (if age is negative).

type person struct {
	name string
	age  int
}

func createPerson(name string, age int) (*person, error) {
	if age<0 {
		return nil, errors.New("Negative age")
	}
	return &person{name:name, age:age}, nil
}

func main() {
	fmt.Println(createPerson("RDJ", 66))
}
