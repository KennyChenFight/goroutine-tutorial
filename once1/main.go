package main

import (
	"bytes"
	"fmt"
	"sync"
)

var once sync.Once

func GetInstanceWithOnce() *singleton {
	once.Do(func() {
		instance = &singleton{buffer: bytes.NewBuffer(make([]byte, 1000))}
	})
	return instance
}

func GetInstance() *singleton {
	return &singleton{buffer: bytes.NewBuffer(make([]byte, 1000))}
}

var instance *singleton

type singleton struct {
	buffer *bytes.Buffer
}

func main() {
	s1 := GetInstance()
	s2 := GetInstanceWithOnce()
	fmt.Println(s1, s2)
}
