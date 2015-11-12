package main

import "fmt"

func main() {
	// slice here is reference type, no copy happened when ranging
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		if i == 0 {
			// change s's size will not deduced the iteration times
			s = s[:3]
			s[2] = 100
		}

		fmt.Println(i, v) // s[2] will changed
	}
	fmt.Println(s) // s[2] will changed
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0 1											  //
// 1 2											  //
// 2 100										  //
// 3 4											  //
// 4 5											  //
// [1 2 100]									  //
////////////////////////////////////////////////////
