package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan)
	//intChan <- 3
	a := <- intChan
	fmt.Println(a)
}
