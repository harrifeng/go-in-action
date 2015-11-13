package main

import "fmt"

func test() (int, int) {
	return 1, 2
}

func main() {
	// Error => multiple-value test() in single-value context
	// s := make([]int, 2)
	// s = test()
	x, _ := test()
	fmt.Println(x)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1											  //
////////////////////////////////////////////////////
