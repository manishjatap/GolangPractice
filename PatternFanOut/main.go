package main

import (
	"fmt"
	"sync"
)

func main() {
	const jobs int = 10
	const workers int = 3

	ch := make(chan int, jobs)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		go func() {
			for {
				select {
				case job := <-ch:
					fmt.Printf("worker %d doing job %d\n", i, job)
					wg.Done()
				}
			}
		}()
	}

	// Pushing jobs to the channel
	for i := 0; i < jobs; i++ {
		wg.Add(1)
		ch <- i
	}

	wg.Wait()
}
