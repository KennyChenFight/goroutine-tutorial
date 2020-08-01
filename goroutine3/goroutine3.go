package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	num := 300000
	wg.Add(num)
	start := time.Now().Unix()
	for i := 1; i <= num; i++ {
		go findPrimes(i, wg)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start, "seconds")
}

func findPrimes(num int, wg *sync.WaitGroup) {
	defer wg.Done()
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
