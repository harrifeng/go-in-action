package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C = "c"
		D        // same with previous
		E = iota // return to interrupted increasing number
		F
	)
	fmt.Println(A, B, C, D, E, F)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0 1 c c 4 5									  //
////////////////////////////////////////////////////
