package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	evenC := make(chan bool)
	oddC := make(chan bool)
	doneC := make(chan bool)

	go even(&wg, evenC, oddC, doneC)
	go odd(&wg, evenC, oddC, doneC)

	evenC <- true
	wg.Wait()
}

func even(wg *sync.WaitGroup, evenC, oddC, doneC chan bool) {
	defer wg.Done()

	for i := 0; i < 10; i += 2 {
		<-evenC
		fmt.Println("even ", i)
		oddC <- true
	}

	doneC <- true
}

func odd(wg *sync.WaitGroup, evenC, oddC, doneC chan bool) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		<-oddC
		fmt.Println("odd ", i)

		select {
		case evenC <- true:
		case <-doneC:
			return
		}

	}
}
