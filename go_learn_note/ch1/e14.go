package main

import "fmt"

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	fmt.Println("Color number is ", c)
}

func main() {
	c := Black
	test(c)

	// Const is auto convert
	test(1)

	x := 1
	_ = x
	// Error! can not use x (type int) as type Color in function argument
	// test(x)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Color number is  0							  //
// Color number is  1							  //
////////////////////////////////////////////////////
