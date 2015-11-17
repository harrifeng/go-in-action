package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	u1 := User{"Tom", 20}
	fmt.Println(u1)

	// u2 := User{"Tom"} // too few values in struct initializer
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// {Tom 20}										  //
////////////////////////////////////////////////////
