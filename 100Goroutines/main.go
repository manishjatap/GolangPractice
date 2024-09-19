package main

import "fmt"

func main() {
	var ch [100]chan int
	output := make(chan int)

	for i := 0; i < 100; i++ {
		ch[i] = make(chan int)
		go print(ch[i], output)
	}

	for i := 0; i < 100; i++ {
		ch[i] <- i
		<-output
	}
}

func print(ch, output chan int) {
	for {
		i := <-ch
		fmt.Println(i)
		output <- i
	}
}
