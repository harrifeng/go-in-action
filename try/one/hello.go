package main

import "fmt"

type User struct {
	id int
}

type Manager struct {
	*User
	title string
}

func main() {

	u := User{12}
	m := &Manager{&u, "hello"}
	fmt.Println(u)
	fmt.Println(m)
	fmt.Println(*m.User)
}
