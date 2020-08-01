package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(2 * time.Second)

	//time.Sleep(2 * time.Second)
	//fmt.Println(timer.Stop())
	//<- timer.C

	//time.Sleep(2 * time.Second)
	//fmt.Println(timer.Reset(10 * time.Second))
	//fmt.Println(<- timer.C)


	timer := time.NewTimer(2 * time.Second)
	for {
		// loading...
		time.Sleep(3 * time.Second)
		if !timer.Stop() {
			select {
			case <- timer.C:
			default:
			}
		}
		timer.Reset(2 * time.Second)
		select {
		case t := <- timer.C:
			fmt.Println("tick at:", t)
		}
	}
}
