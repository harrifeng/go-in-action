package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func main() {

	m := Manager{User{1, "Tom"}, "Administrator"}
	fmt.Println(m)
	// cannot use m (type Manager) as type User in assignment,
	// no Polymorphism here
	// var u User = m

}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// {{1 Tom} Administrator}						  //
////////////////////////////////////////////////////
