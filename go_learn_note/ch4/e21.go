package main

import "fmt"

type User struct {
	name string
}

type Manager struct {
	User
	title string
}

func main() {

	m := Manager{
		User:  User{"Tom"},
		title: "Administrator",
	}

	fmt.Println(m)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// {{Tom} Administrator}						  //
////////////////////////////////////////////////////
