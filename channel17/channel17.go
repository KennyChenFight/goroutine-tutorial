package main

import (
	"fmt"
	"time"
)

func main() {
	//quit := make(chan bool, 1)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	quit <- true
	//}()
	//
	//select {
	//case <- quit:
	//	fmt.Println("over")
	//case <- time.After(2 * time.Second):
	//	fmt.Println("timeout")
	//}

	for {
		select {
		case t := <- time.After(2 * time.Second):
			fmt.Println("tick at:", t)
		}
	}
}
