package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	new_f := squares()
	fmt.Println(new_f())
	fmt.Println(new_f())
	fmt.Println(squares()())
	fmt.Println(squares()())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1											  //
// 4											  //
// 9											  //
// 16											  //
// 1											  //
// 4											  //
// 1											  //
// 1											  //
////////////////////////////////////////////////////
