package main

import "fmt"

func main() {
	//var intChan chan int
	////fmt.Println(intChan)
	////intChan <- 10
	//
	//intChan = make(chan int, 1)
	//fmt.Println(cap(intChan))
	//fmt.Println(len(intChan))
	//intChan <- 10
	//fmt.Println(len(intChan))
	//fmt.Println(<- intChan)
	//fmt.Println(len(intChan))

	//var intChen = make(chan int, 3)
	//intChen <- 1
	//intChen <- 2
	//intChen <- 3
	//b := <- intChen
	//intChen <- 4
	//fmt.Println(b)

	stringChan := make(chan string)
	go func() {
		stringChan <- "hello world"
	}()
	fmt.Println(<- stringChan)
}
