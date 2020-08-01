package main

import (
	"fmt"
	"time"
)

func main() {
	num := 300000
	for i := 1; i <= num; i++ {
		go findPrimes(i)
	}
	time.Sleep(5 * time.Second)
}

func findPrimes(num int) {
	if num == 1 {
		return
	} else if num == 2 {
		fmt.Println(num)
	} else {
		for i := 2; i < num; i++ {
			if num % i == 0 {
				return
			}
		}
		fmt.Println(num)
	}
}
