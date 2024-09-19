package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := range ch {
			fmt.Printf("Executing Job %d\n", i)
		}
	}()

	var jobs = 30

	for i := 0; i < jobs; i++ {
		select {
		case ch <- i:
		default:
			fmt.Printf("Dropping a Job %d\n", i)
		}
	}
}
