package main

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0 1 2 3 4 5 6								  //
////////////////////////////////////////////////////
