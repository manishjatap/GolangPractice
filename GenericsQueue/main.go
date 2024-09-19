package main

import "fmt"

// type QueueInteface interface {
// 	EnQueue()
// 	DeQueue()
// }

// type Queue[T any] struct {

// }

// func main() {

// }

type sss interface {
}

type data struct {
	a sss
}

type p struct {
	b int
}

func main() {
	pp := p{b: 10}

	mymap := map[data]int{
		{a: pp}: 1,
		{a: pp}: 1,
	}

	fmt.Println(mymap)
}
