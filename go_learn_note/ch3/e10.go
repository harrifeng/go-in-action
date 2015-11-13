package main

import "fmt"

func main() {
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello World" }
	fmt.Println((<-fc)())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Hello World									  //
////////////////////////////////////////////////////
