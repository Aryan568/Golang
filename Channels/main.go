package main

import "fmt"

func main() {
	fmt.Println("Channels")

	num:= make(chan int)
	str:= make(chan string)

	go data(num, str)

	fmt.Println("Channel Number",<-num)
	fmt.Println("Channel String",<-str)
}

func data(num chan int, str chan string) {
	num<-99
	str<-"Golang"
}