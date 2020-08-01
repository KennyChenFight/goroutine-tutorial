package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func hello2() {
	fmt.Println("hello2")
	c <- true
}

var c = make(chan bool, 1)

func main() {
	// time.AfterFunc(2 * time.Second, hello)
	//
	//time.Sleep(3 * time.Second)
	//fmt.Println("over")

	time.AfterFunc(2 * time.Second, hello2)
	<-c
	fmt.Println("over timer")
}
