package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go task(ctx)
	//time.Sleep(3 * time.Second)
	//cancel()
	//time.Sleep(3 * time.Second)


	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-shutdown:
		fmt.Println("shutdown")
		cancel()
		time.Sleep(2 * time.Second)
	}
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
