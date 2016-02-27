package main

import (
	"fmt"
)

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	c3 = iota
	c4
	c5
)

func main() {
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0											  //
// 1											  //
// 2											  //
// 0											  //
// 1											  //
// 2											  //
////////////////////////////////////////////////////
