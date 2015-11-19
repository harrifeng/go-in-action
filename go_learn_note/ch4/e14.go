package main

import "fmt"

func main() {
	data1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data1[:]
	fmt.Println(s1)
	data2 := [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s2 := data2[:]
	fmt.Println(s2)

	copy(s1, s2) // copy s1 from s2

	fmt.Println(s1)
	fmt.Println(s2)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [0 1 2 3 4 5 6 7 8 9]						  //
// [10 11 12 13 14 15 16 17 18 19 20]			  //
// [10 11 12 13 14 15 16 17 18 19]				  //
// [10 11 12 13 14 15 16 17 18 19 20]			  //
////////////////////////////////////////////////////