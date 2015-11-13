package main

import "fmt"

func main() {
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World" },
	}

	fmt.Println(d.fn())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Hello, World									  //
////////////////////////////////////////////////////
