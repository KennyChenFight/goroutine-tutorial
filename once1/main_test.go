package main

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstance()
	}
}

func Benchmark_withOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceWithOnce()
	}
}
