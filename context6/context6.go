package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "username", "kenny")
	go task(ctx)
	time.Sleep(3 * time.Second)
}

func task(ctx context.Context) {
	fmt.Println("task get value:", ctx.Value("username"))
}

