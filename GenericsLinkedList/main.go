package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (node *Node[any]) Add(key any) {
	if node == nil {
		node = &Node[any]{
			Value: key,
			Next:  nil,
		}
		return
	}

	temp := node
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = &Node[any]{
		Value: key,
		Next:  nil,
	}
}

func (node *Node[any]) Remove() {

}

func (node *Node[any]) Print() {
	temp := node
	for temp != nil {
		fmt.Print(temp.Value, " ")
		temp = temp.Next
	}

	fmt.Println()
}

func main() {
	list := Node[any]{}
	list.Add("hi")
	list.Add(1)
	list.Add(2)
	list.Print()
}
