package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		task1()
		cancel1()
	}()
	go func() {
		task2()
		cancel2()
	}()

	<- ctx1.Done()
	<- ctx2.Done()
}

func task1() {
	fmt.Println("task1 work...")
	time.Sleep(3 * time.Second)
	fmt.Println("task1 over...")
}

func task2() {
	fmt.Println("task2 work...")
	time.Sleep(2 * time.Second)
	fmt.Println("task2 over...")
}
