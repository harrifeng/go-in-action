package main

import "fmt"

var globalX int
var globalY = "abc"

func main() {
	fmt.Println(globalX)
	fmt.Println(globalY)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0											  //
// abc											  //
////////////////////////////////////////////////////
