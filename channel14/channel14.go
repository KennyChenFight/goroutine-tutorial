package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		number := rand.Intn(10) + 1
		result <- number
	}()

	select {
	case number := <- result:
		fmt.Println("get number:", number)
	case <- time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
