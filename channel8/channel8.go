package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(intChan chan int) {
	count := 0
	for {
		if count != 6 {
			num := rand.Intn(10) + 1
			intChan <- num
			fmt.Println("put num:", num)
			count++
		} else {
			close(intChan)
			break
		}
	}
}

func main() {
	intChan := make(chan int, 1)
	go generateNumbers(intChan)
	for num := range intChan {
		fmt.Println("get num:", num)
		time.Sleep(time.Second)
	}
}
