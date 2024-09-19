package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	ch := make(chan any, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		arr := []any{"asdfds", 12, 4.3, true, 's'}
		for i := range 5 {
			fmt.Println(reflect.TypeOf(arr[i]))
			ch <- arr[i]
		}
		close(ch)
	}()

	go func() {
		for j := range ch {
			fmt.Println(reflect.TypeOf(j), j)
		}
		wg.Done()
	}()

	wg.Wait()
}
