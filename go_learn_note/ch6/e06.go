package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
	var o interface{} = &User{1, "Tom"}

	//////////////////////////////////
	// Source Code for golang	    //
	// type Stringer interface {    //
	// 	String() string			    //
	// }						    //
	//////////////////////////////////

	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}

	u := o.(*User)

	// u := o.(User)  // panic: interface conversion: interface is *main.User, not main.User
	fmt.Println(u)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1, Tom										  //
// 1, Tom										  //
////////////////////////////////////////////////////
