package main

import "fmt"

func main() {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	fmt.Println(m[1].name)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// user1										  //
////////////////////////////////////////////////////
