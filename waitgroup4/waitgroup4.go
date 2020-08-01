package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var boss sync.WaitGroup
var supervisor1 sync.WaitGroup
var supervisor2 sync.WaitGroup

func work1() {
	defer supervisor1.Done()
	workSecond := rand.Intn(10) + 1
	time.Sleep(time.Duration(workSecond) * time.Second)
	fmt.Println("one work1 done")
}

func work2() {
	defer supervisor2.Done()
	workSecond := rand.Intn(10) + 1
	time.Sleep(time.Duration(workSecond) * time.Second)
	fmt.Println("one work2 done")
}

func wait1() {
	defer boss.Done()
	supervisor1.Wait()
	fmt.Println("all work1 done")
}

func wait2() {
	defer boss.Done()
	supervisor2.Wait()
	fmt.Println("all work2 done")
}

func waitAllWork() {
	go wait1()
	go wait2()
	boss.Wait()
	fmt.Println("all work done")
}

func main() {
	supervisor1.Add(3)
	for i := 1; i <= 3; i++ {
		go work1()
	}

	supervisor2.Add(3)
	for i := 1; i <= 3; i++ {
		go work2()
	}

	boss.Add(2)
	waitAllWork()

	fmt.Println("let's celebrate!")
}
