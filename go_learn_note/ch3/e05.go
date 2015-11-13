package main

import "fmt"

func test() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(test()))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 3											  //
////////////////////////////////////////////////////
