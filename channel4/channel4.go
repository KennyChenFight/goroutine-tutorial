package main

func main() {
	intChan := make(chan <- int, 3)
	stringChan := make(<- chan int, 3)
	intChan <- 1
	a := <- stringChan
	stringChan <- "hello"
}
