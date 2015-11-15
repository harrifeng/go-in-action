package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 8: 100} // Note! there are nothing in []
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // make to create, len and cap as parameter
	fmt.Println(s2, len(s1), cap(s1))

	s3 := make([]int, 6) // len == cap here
	fmt.Println(s3, len(s3), cap(s3))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [0 1 2 3 0 0 0 0 100] 9 9					  //
// [0 0 0 0 0 0] 9 9							  //
// [0 0 0 0 0 0] 6 6							  //
////////////////////////////////////////////////////
