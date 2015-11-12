package main

import "fmt"

func main() {
	// array is NOT reference type, everytime range will
	// copy the value to i and v
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a) // change a's value succesfully
		}
		a[i] = v + 100 // use a's copied value to construct the final version

	}
	fmt.Println(a)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [0 999 999]									  //
// [100 101 102]								  //
////////////////////////////////////////////////////
