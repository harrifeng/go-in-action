package main

import "fmt"

func main() {
	a := []int{0, 0, 0}
	a[1] = 10

	fmt.Println(a)

	// make return object
	b := make([]int, 3)
	b[1] = 10

	fmt.Println(b)

	// new return the pointer
	c := new([]int)
	fmt.Println(*c)
}
