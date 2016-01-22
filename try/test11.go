package main

import "fmt"

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {

	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))

	fmt.Printf("%T\n", f)

	// cannot use product (type func(int, int) int) as type func(int) int in assignment
	// f = product

}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 9											  //
// -3											  //
// func(int) int								  //
////////////////////////////////////////////////////
