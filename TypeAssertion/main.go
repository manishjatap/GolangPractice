package main

import "fmt"

type Shape interface {
	Area()
} // Empty interface

type Circle struct {
	Color1 string
}

type Rectangle struct {
	Color2 string
}

func (r *Rectangle) Area() {

	fmt.Println(r.Color2)

}
func (c *Circle) Area() {

	fmt.Println(c.Color1)

}

func main() {
	var s1, s2 Shape
	s1 = Circle{Color1: "red"}
	s2 = Rectangle{Color2: "black"}
	s1.Area()
	s2.Area()
	Area(c)
	Area(r)
}

// type Shape interface {
// } // Empty interface

// type Circle struct {
// 	Color1 string
// }

// type Rectangle struct {
// 	Color2 string
// }

// func Area(s Shape) {
// 	switch ss := s.(type) {
// 	case Circle:
// 		fmt.Println(ss.Color1)
// 	case Rectangle:
// 		fmt.Println(ss.Color2)
// 	}
// }

// func main() {
// 	c := Circle{Color1: "red"}
// 	r := Rectangle{Color2: "black"}
// 	Area(c)
// 	Area(r)
// }
