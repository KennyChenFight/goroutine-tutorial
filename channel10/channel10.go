package main

import "fmt"

func main() {
	size := 300000
	numbers := make(chan int, size)
	primeNumbers := make(chan int, size)
	done := make(chan bool, size)

	go putNumbers(size, numbers)
	go findPrimes(numbers, primeNumbers, done)

	for primeNumber := range primeNumbers {
		fmt.Println(primeNumber)
	}
}

func putNumbers(size int, numbers chan int) {
	for i := 1; i <= size; i++ {
		numbers <- i
	}
	close(numbers)
}

func findPrimes(numbers chan int, primeNumbers chan int, done chan bool) {
	for number := range numbers {
		go checkPrime(number, primeNumbers, done)
	}
	for i := 0; i < cap(done); i++ {
		<- done
	}
	close(primeNumbers)
}

func checkPrime(num int, primeNumbers chan int, done chan bool) {
	defer func() {
		done <- true
	}()
	if num == 1 {
		return
	} else if num == 2 {
		primeNumbers <- num
	} else {
		for i := 2; i < num; i++ {
			if num % i == 0 {
				return
			}
		}
		primeNumbers <- num
	}
}
