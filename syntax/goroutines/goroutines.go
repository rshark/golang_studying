package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("in closure before send channel")
		c <- 1
		fmt.Println("in closure after send channel")
	}()
	fmt.Println("out closure before receive channel")
	<-c
	fmt.Println("out closure after receive channel")
}
