package main

import "fmt"

func test() (int, int) {
	return 1, 2
}

func sum(n ...int) int {
	var x int

	for _, i := range n {
		x += i
	}

	return x
}

func main() {
	fmt.Println(sum(test()))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 3											  //
////////////////////////////////////////////////////
