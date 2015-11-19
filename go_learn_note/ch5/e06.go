package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self User) Test() {
	fmt.Printf("%p, %v\n", &self, self)
}

func main() {
	u := User{2, "Jacky"}
	u.Test()

	mValue := u.Test
	mValue()

	mExpression := (User).Test
	mExpression(u)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0x8201e41a0, {2 Jacky}						  //
// 0x8201e4200, {2 Jacky}						  //
// 0x8201e4240, {2 Jacky}						  //
////////////////////////////////////////////////////
