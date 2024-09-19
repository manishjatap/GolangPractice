package main

import (
	"fmt"
	"stack"
)

func main() {
	tree := []int{1, 2, 3, 4, 5, 6}
	inorder(tree)
}

func inorder(tree []int) {
	var s stack.Stack

	fmt.Println(s)
	// mymap := make(map[int]int)
	// size := len(tree)

	// s.Push(0) // initial push
	// for {
	// 	node := s.Pop()
	// 	right := (node * 2) + 1
	// 	left := (node * 2) + 2

	// 	if _, found := mymap[node]; !found {
	// 		s.Push(left)
	// 		s.Push(node)
	// 		s.Push(right)
	// 	} else {
	// 		mymap[node] = node
	// 	}

	// }
}
