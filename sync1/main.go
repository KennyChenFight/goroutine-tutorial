package main

import (
	"sync"
)

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

type singleton struct {
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {

			GetInstance()
		}()
	}
}
