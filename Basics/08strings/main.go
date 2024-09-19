package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`Strings`) //backticks(``) are also used to represent strings

	name:= `Golang`
// access a particular character of a string
	fmt.Printf("%c\n", name[0])
	fmt.Printf("%c\n", name[1])
	fmt.Printf("%c\n", name[2])
	fmt.Printf("%c\n", name[3])
	fmt.Printf("%c\n", name[4])
	fmt.Printf("%c\n", name[5])
	

	fmt.Println("Length of the string:", len(name))  //string length

	name1:=`Go`
	name2:=`lang`

	fmt.Println(name1+name2) // string concatenate

	name3:= `Golang`

	fmt.Println(strings.Compare(name, name1))  //1
	fmt.Println(strings.Compare(name1, name3))	//-1
	fmt.Println(strings.Compare(name3, name))	//0

	text:= `Go Programming`

	result:= strings.Contains(text, name1)  // check if the text(string) contains name1(substring)
	fmt.Println(result)

	result= strings.Contains(text, name3)
	fmt.Println(result)

	name4:= `cost`

	fmt.Println("Old String", name4)

	replacedChr:= strings.Replace(name4, `c`, `m`, 1)    //replace character from string

	fmt.Println(`New String:`,replacedChr)

	fmt.Println(strings.ToUpper(name4))     //upper case
	fmt.Println(strings.ToLower(text))		//lower case

	texts:= "Gem:ini"
	parts:= strings.Split(texts, ":")
	fmt.Println(parts[0]+"\n"+parts[1])
}