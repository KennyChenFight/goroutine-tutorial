package main

import "fmt"

func main() {
	num := 100
	intChan := make(chan int, 100)
	for i := 0; i < num; i++ {
		intChan <- i
	}
	close(intChan)

	for i := range intChan {
		fmt.Println(i)
	}
}
