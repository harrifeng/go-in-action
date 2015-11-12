package main

import "fmt"

func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {

	fmt.Println(test("sum: %d", 1, 2, 3))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(test("sum: %d", s...))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// sum: 6										  //
// sum: 15   									  //
////////////////////////////////////////////////////
