package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	duration := 1 * time.Second
	ctx, delay := context.WithTimeout(context.Background(), duration)
	defer delay() // need to check why this is needed?
	waitForResult := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		waitForResult <- 1
	}()

	select {
	case <-waitForResult:
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
