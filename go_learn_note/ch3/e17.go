package main

import "fmt"

func test() {
	defer recover()
	defer fmt.Println(recover())

	defer func() {
		func() {
			fmt.Println("defer inner")
			recover()
		}()
	}()

	panic("test panic")
}

func main() {
	test()
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// defer inner									  //
// <nil>										  //
// panic: test panic							  //
////////////////////////////////////////////////////
