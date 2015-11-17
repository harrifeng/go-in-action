package main

import "fmt"

func main() {
	type user struct{ name string }

	m := map[int]user{
		1: {"user1"},
	}

	fmt.Println(m)

	// m[1].name = "Tom"			// cannot assign to m[1].name
	u := m[1]
	u.name = "Tom"
	m[1] = u
	fmt.Println(m)

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	fmt.Println(m2[1])

	m2[1].name = "Jack"

	fmt.Println(m2[1])
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// map[1:{user1}]								  //
// map[1:{Tom}]									  //
// &{user1}										  //
// &{Jack}										  //
////////////////////////////////////////////////////
