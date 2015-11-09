package main

import "fmt"

func main() {
	a := 0x011
	b := 0x101
	fmt.Println(a & b)
	fmt.Println(0x0001)

	fmt.Println(a | b)
	fmt.Println(0x0111)

	fmt.Println(a ^ b)
	fmt.Println(0x0110)

	// clear 1 to 0 on 0th and 3th from 0x101 for 0x011
	fmt.Println(a &^ b)
	fmt.Println(0x010)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1											  //
// 1											  //
// 273											  //
// 273											  //
// 272											  //
// 272											  //
// 16											  //
// 16											  //
////////////////////////////////////////////////////
