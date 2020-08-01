package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := make(chan int, 1)
	quit := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		number := rand.Intn(10) + 1
		numbers <- number

		// close resource
		time.Sleep(3 * time.Second)
		quit <- true
	}()

	for {
		select {
		case number := <-numbers:
			fmt.Println("get number:", number)
			// handle
		case <- quit:
			fmt.Println("over~")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(time.Second)
		}
	}
}

