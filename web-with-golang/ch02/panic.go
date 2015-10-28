package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER2")

func test() {
	if user == "" {
		panic("no value for $USER")
	}
	fmt.Println(user)
}

func main() {
	test()
}

//////////////////////////////////////////////////////////
// <===================OUTPUT===================>	    //
// panic: no value for $USER						    //
// 													    //
// goroutine 16 [running]:							    //
// runtime.panic(0x495080, 0xc208000150)			    //
// 	/usr/lib/go/src/pkg/runtime/panic.c:279 +0xf5	    //
// ...												    //
//////////////////////////////////////////////////////////
