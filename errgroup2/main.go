package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				log.Printf("goroutine should cancel")
				return nil
			default:
				if err := getPage("https:kenny.example.com"); err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			}
		}
		return nil
	})
	eg.Go(func() error {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				log.Printf("goroutine should cancel")
				return nil
			default:
				if err := getPage("https://google.com"); err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			}
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("get error: %v", err)
	}
}

func getPage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to get page: %s, wrong statusCode: %d", url, resp.StatusCode)
	}
	log.Printf("success get page %s", url)
	return nil
}
