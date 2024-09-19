package main

import (
	"fmt"
)

func main() {
	print(0)
}

func print(i int) {
	defer func() {
		if r := recover(); r != nil {
			// debug.PrintStack()
			fmt.Println("Error: Invalid input")
		}
	}()

	fmt.Println(1 / i)
}
