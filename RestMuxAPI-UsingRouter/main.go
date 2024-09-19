package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	a := mux.NewRouter()
	fmt.Println(a)
}
