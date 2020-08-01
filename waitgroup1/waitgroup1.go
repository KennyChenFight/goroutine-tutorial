package main

import "sync"

func main() {
	var wg1 sync.WaitGroup
	wg2 := new(sync.WaitGroup)
	wg3 := &sync.WaitGroup{}
	var wg4 = &sync.WaitGroup{}
}
