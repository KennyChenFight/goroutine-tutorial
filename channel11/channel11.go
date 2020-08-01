package main

import (
	"fmt"
	"math/rand"
	"time"
)

type worker struct {
	id int
	number int
}

func work(jobs <- chan worker, results chan <-worker) {
	for j := range jobs {
		fmt.Println("worker", j.id, "started job")
		time.Sleep(time.Duration(rand.Intn(10) + 1) * time.Second)
		fmt.Println("worker", j.id, "finished job")
		j.number = rand.Intn(10) + 1
		results <- j
	}
}

func main() {
	number := 10
	jobs := make(chan worker, number)
	results := make(chan worker, number)

	for i := 1; i <= 5; i++ {
		go work(jobs, results)
	}

	for j := 1; j <= number; j++ {
		worker := worker{id: j}
		jobs <- worker
	}
	close(jobs)

	for a := 1; a <= number; a++ {
		result := <- results
		fmt.Printf("workerId: %d, result: %v\n", result.id, result.number)
	}
}

