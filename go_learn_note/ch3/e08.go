package main

import "fmt"

func main() {

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}

	fmt.Println(fns[0](100))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 101											  //
////////////////////////////////////////////////////
