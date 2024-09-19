package main

import "fmt"

type Stack[y any] struct {
	Elem []any
}

func (stack *Stack[any]) Push(key any) {
	fmt.Printf(" Pushing %v to Stack \n", key)
	stack.Elem = append(stack.Elem, key)
}

func (stack *Stack[any]) Pop() {
	if len(stack.Elem) < 1 {
		return // ignore pop operation
	}
	fmt.Printf(" Popped %v from Stack \n", stack.Elem[len(stack.Elem)-1])
	stack.Elem = stack.Elem[:len(stack.Elem)-1]
}

func (stack *Stack[any]) Print() {
	for i := 0; i < len(stack.Elem); i++ {
		fmt.Print(stack.Elem[i], " ")
	}
	fmt.Println()
}

func main() {
	stack := Stack[int]{nil}
	stack.Push(1)
	stack.Push(3)
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
}
