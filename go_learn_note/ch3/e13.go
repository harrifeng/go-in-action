package main

import "fmt"

func test(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x)
	}()

	defer fmt.Println("c")
}

func main() {
	test(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// c											  //
// b											  //
// a											  //
// panic: runtime error: integer divide by zero	  //
////////////////////////////////////////////////////
