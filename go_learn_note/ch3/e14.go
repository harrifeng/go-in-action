package main

import "fmt"

func test() {
	x, y := 10, 20

	// take x as parameter for anoymous function, saved x's value
	// as 10.
	defer func(i int) {
		fmt.Println("defer:", i, y)
	}(x)

	x += 10
	y += 100
	fmt.Println("x = ", x, "y = ", y)
}

func main() {
	test()
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// x =  20 y =  120								  //
// defer: 10 120								  //
////////////////////////////////////////////////////
