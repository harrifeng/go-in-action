package main

import "fmt"

func main() {
	data, i := [3]int{0, 1, 2}, 0
	fmt.Println(i)
	fmt.Println(data)

	// When assigning multiple values in one line
	// You should:
	// [1] make sure all the value: (i => 0) and (data[i] => data[0])
	// [2] assign value base on [1]: i = 2, data[0] = 100
	i, data[i] = 2, 100
	fmt.Println(i)
	fmt.Println(data)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0											  //
// [0 1 2]										  //
// 2											  //
// [100 1 2]									  //
////////////////////////////////////////////////////
