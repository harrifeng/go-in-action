package main

import "fmt"

type Resource struct {
	id   int
	name string
}

type Classify struct {
	id int
}

type User struct {
	Resource
	Classify
	name string
}

func main() {
	u := User{
		Resource{1, "people"},
		Classify{100},
		"Jack",
	}

	fmt.Println(u.name)
	fmt.Println(u.Resource.name)
	// fmt.Println(u.id)			//  ambiguous selector u.id
	fmt.Println(u.Classify.id)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Jack											  //
// people										  //
// 100											  //
////////////////////////////////////////////////////
