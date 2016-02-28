package main

import "fmt"

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg)
	}
}

func myfuncactually(args []int) {
	for _, arg := range args {
		fmt.Print(arg)
	}
}

func main() {
	myfunc(1, 2, 3)
	fmt.Println()
	myfunc()
	fmt.Println()
	myfuncactually([]int{4, 5, 6})
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 123											  //
// 												  //
// 456											  //
////////////////////////////////////////////////////
