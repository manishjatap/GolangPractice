package main

import "fmt"

type Stack struct {
	Data []int
}

func (s *Stack) Push(item int) {
	if s.Data == nil {
		s.Data = []int{}
	}
	s.Data = append(s.Data, item)
}

func (s *Stack) Pop() int {
	if len(s.Data) == 0 {
		return -1
	}
	item := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return item
}

func (s *Stack) Print() {
	for _, val := range s.Data {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
