package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3 * time.Second))
	go task(ctx)
	time.Sleep(5 * time.Second)
	//time.Sleep(1 * time.Second)
	//cancel()
	//time.Sleep(1 * time.Second)
}

func task(ctx context.Context) {
	fmt.Println("task work...")
	go subTask("one", ctx)
	go subTask("two", ctx)
	go subTask("three", ctx)
}

func subTask(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("subTask %s stop work by ctx cancel, err: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("subTask %s work...\n", name)
			time.Sleep(time.Second)
		}
	}
}
