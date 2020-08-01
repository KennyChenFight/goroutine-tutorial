package main

import (
	"fmt"
	"time"
)

func main() {
	num := 100
	intChan := make(chan int, num)

	//for i := 0; i < num; i++ {
	//	intChan <- i
	//}
	//
	//for i := 0; i < num; i++ {
	//	fmt.Println(<- intChan)
	//}


	go func() {
		for i := 0; i < 200; i++ {
			intChan <- i
		}
	}()

	go func() {
		for i := 0; i < num; i++ {
			fmt.Println(<- intChan)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("剩餘:", len(intChan))
}
