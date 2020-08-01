package main

import "fmt"

type dog struct {
	name string
	gender string
}

func main() {
	allChan := make(chan interface{}, 3)
	allChan <- dog{"jack", "male"}
	allChan <- 1
	allChan <- "hello World"

	d := <- allChan
	fmt.Println(d)
}
