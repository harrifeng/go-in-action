package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(b)
	)
	fmt.Println(a, b, c)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// abc 3 8										  //
////////////////////////////////////////////////////
