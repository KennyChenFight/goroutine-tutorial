package main

import (
	"fmt"
	"time"
)

func produce(numbers chan <- int) {
	for i := 0; i < 5; i++ {
		//time.Sleep(3 * time.Second)
		numbers <- i
	}
}

func consume(numbers <- chan int) {
	for {
		select {
		case number := <- numbers:
			fmt.Println("consume success:", number)
		case <- time.After(2 * time.Second):
			fmt.Println("timeout! continue...")
		}
	}
}

func main() {
	numbers := make(chan int)
	go produce(numbers)
	go consume(numbers)

	select {

	}
}
