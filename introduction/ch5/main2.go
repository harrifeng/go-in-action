package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1 odd										  //
// 2 even										  //
// 3 odd										  //
// 4 even										  //
// 5 odd										  //
// 6 even										  //
// 7 odd										  //
// 8 even										  //
// 9 odd										  //
// 10 even										  //
////////////////////////////////////////////////////
