package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	luckyNumber := 10
	go guess(ctx, luckyNumber)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func guess(ctx context.Context, luckyNumber int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx cancel, can not guess right number")
			return
		default:
			if num := rand.Intn(5) + 1; num == luckyNumber {
				fmt.Printf("guess luckyNumber: %d\n", num)
				return
			} else {
				fmt.Printf("guess wrong number: %d\n", num)
			}
		}
	}
}
