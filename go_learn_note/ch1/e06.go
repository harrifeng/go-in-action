package main

import "fmt"

func test() (int, string) {
	return 1, "abc"
}

func main() {
	_, s := test()
	fmt.Println(s)
	i := 0
	_ = i
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// o											  //
// abc											  //
////////////////////////////////////////////////////
