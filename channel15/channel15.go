package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	<- timer.C
	fmt.Println("timeout1")

	timer.Reset(2 * time.Second)
	<- timer.C
	fmt.Println("timeout2")
}
