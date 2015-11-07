package main

import "fmt"

func main() {
	const (
		_        = iota             // itoa = 0
		KB int64 = 1 << (10 * iota) // itoa = 1
		MB                          // Same with KB, but itoa = 2 here
		GB
		TB
	)
	fmt.Println(KB, MB, GB, TB)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1024 1048576 1073741824 1099511627776		  //
////////////////////////////////////////////////////
