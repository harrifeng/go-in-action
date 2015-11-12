package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	i := 2

	switch i {
	case x[1]:
		fmt.Println("a")
	case 1, 3:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// a											  //
////////////////////////////////////////////////////
