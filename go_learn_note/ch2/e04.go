package main

import "fmt"

func main() {
	// string/array/slice they are the same
	s := "abc"
	for i, j := range s {
		fmt.Printf("%d  %c\n", i, j)
	}

	for _, c := range s {
		fmt.Println(c)
	}

	//map
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0  a											  //
// 1  b											  //
// 2  c											  //
// 97											  //
// 98											  //
// 99											  //
// a 1											  //
// b 2											  //
////////////////////////////////////////////////////
