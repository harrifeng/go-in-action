package main

import "fmt"

func main() {
	x := 1234
	type bigint int64
	var b bigint = bigint(x)
	var b2 int64 = int64(b)

	type myslice []int
	var s myslice = []int{1, 2, 3}
	var s2 []int = s

	fmt.Println(b2)
	fmt.Println(s2)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1234											  //
// [1 2 3]										  //
////////////////////////////////////////////////////
