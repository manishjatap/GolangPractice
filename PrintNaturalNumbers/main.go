package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid input")
		return
	}
	inputStr := os.Args[1]
	_, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("invalid input")
		return
	}

}
