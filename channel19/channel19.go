package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at:", t)
		}
	}()
	time.Sleep(2 * time.Second)
	ticker.Stop()
	//time.Sleep(3 * time.Second)
	//ticker.Stop()
	////ticker.Stop()
	//fmt.Println("over")
}
