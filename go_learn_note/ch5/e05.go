package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue()

	mExpression := (*User).Test
	mExpression(&u)

}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0x8201e41a0, &{1 Tom}						  //
// 0x8201e41a0, &{1 Tom}						  //
// 0x8201e41a0, &{1 Tom}						  //
////////////////////////////////////////////////////
