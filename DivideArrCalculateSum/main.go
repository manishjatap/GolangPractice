package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 2, 2, 3, 3, 3, 5, 6, 5, 6}
	n := 2
	var wg sync.WaitGroup

	size := len(arr) / n
	counter := 0

	ch := make(chan int, n)

	for counter < n {
		start := counter * size
		end := start + size

		wg.Add(1)
		if n-counter == 1 {
			go calculateSum(&wg, arr[start:], ch)
		} else {
			go calculateSum(&wg, arr[start:end], ch)
		}
		counter++
	}

	wg.Wait()

	close(ch)
	final := 0
	for i := range ch {
		final += i
	}

	fmt.Println(final)
}

func calculateSum(wg *sync.WaitGroup, arr []int, ch chan int) {
	defer wg.Done()
	sum := 0
	fmt.Println(arr)
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
	ch <- sum
}
