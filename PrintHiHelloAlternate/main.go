package main

import "fmt"

func main() {

	hiC := make(chan bool)
	helloC := make(chan bool)
	doneC := make(chan bool)

	// Print Hi
	go func() {
		for i := 0; i < 10; i++ {
			<-hiC
			fmt.Println("Hi")
			helloC <- true
		}
		doneC <- true
	}()

	// Print Hello
	go func() {
		for i := 0; ; i++ {
			<-helloC
			fmt.Println("Hello")
			hiC <- true
		}
	}()

	hiC <- true
	<-doneC
}
