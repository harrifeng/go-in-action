package main

import "fmt"

func main() {
	x := 10
	switch x {
	case 10:
		fmt.Println("a")
		fallthrough
	case 0:
		fmt.Println("b")
	case -1:
		fmt.Println("c")
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// a											  //
// b											  //
////////////////////////////////////////////////////
