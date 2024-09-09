package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutine-Concurrency")

	go display("Process 1")
	display("Processess")
	go display("Process 2")
	go display("Process 3")

	time.Sleep(time.Second*1)
}

func display(s string) {
	fmt.Println(s)
}