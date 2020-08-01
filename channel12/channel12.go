package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//quit := make(chan bool, 1)
	//go func() {
	//	time.Sleep(time.Second)
	//	quit <- true
	//}()
	//
	//for {
	//	flag := false
	//	select {
	//	case <- quit:
	//		fmt.Println("done")
	//		flag = true
	//	default:
	//		fmt.Println("wait")
	//		time.Sleep(time.Second)
	//	}
	//	if flag {
	//		break
	//	}
	//}
	//fmt.Println("hello")

	//numbers := make(chan int, 5)
	//go func() {
	//	for number := range numbers {
	//		fmt.Println("start consume:", number)
	//		time.Sleep(time.Duration(rand.Intn(10) + 1) * time.Second)
	//		fmt.Println("finish consume:", number)
	//	}
	//}()
	//
	//for {
	//	number := rand.Intn(10) + 1
	//	select {
	//	case numbers <- number:
	//		fmt.Println("produce number:", number)
	//	}
	//}

	task1Chan := make(chan bool, 1)
	task2Chan := make(chan bool, 1)

	go func() {
		seconds := time.Duration(rand.Intn(10) + 1) * time.Second
		time.Sleep(seconds)
		task1Chan <- true
	}()

	go func() {
		seconds := time.Duration(rand.Intn(10) + 1) * time.Second
		time.Sleep(seconds)
		task2Chan <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		case <- task1Chan:
			fmt.Println("task1 over")
		case <- task2Chan:
			fmt.Println("task2 over")
		}
	}
}