package main

import "fmt"

func main() {
	s := "Hello, World!"

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[1:5])
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Hello										  //
// World!										  //
// ello											  //
////////////////////////////////////////////////////
