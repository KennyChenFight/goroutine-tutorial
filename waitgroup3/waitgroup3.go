package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	jobs := 100
	wg.Add(jobs)
	for i := 1; i <= 100; i++ {
		go doTask(wg)
	}
	wg.Wait()
	fmt.Println("jobs all done!")
}

func doTask(wg *sync.WaitGroup) {
	defer wg.Done()
}
