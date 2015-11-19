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

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}, "Administrator"}

	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Manager: 0x8201ec0c0, &{{1 Tom} Administrator} //
// User: 0x8201ec0c0, &{1 Tom}					  //
////////////////////////////////////////////////////
