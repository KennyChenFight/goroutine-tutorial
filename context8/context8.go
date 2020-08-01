package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	ctx := req.Context()
	ctx, _ = context.WithTimeout(ctx, 10 * time.Second)
	ctx = context.WithValue(ctx, "username", "kenny")
	message := make(chan string)
	go sayHello(ctx, message)

	select {
	case <-ctx.Done():
		fmt.Println("server get err:", ctx.Err())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	case m := <- message:
		w.Write([]byte(m))
	}
}

func sayHello(ctx context.Context, message chan <- string) {
	select {
	case <-ctx.Done():
		fmt.Println("sayHello shutdown")
	case <-time.After(5 * time.Second):
		message <- fmt.Sprintf("hello %s", ctx.Value("username"))
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
