package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	jobs := 100
	wg.Add(jobs)
	for i := 1; i <= 100; i++ {
		go doTask(wg)
	}
	wg.Wait()
	fmt.Println("jobs all done!")
}

func doTask(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("one task done")
		wg.Done()
	}()
	number := rand.Intn(10) + 1
	time.Sleep(time.Duration(number) * time.Second)
}
