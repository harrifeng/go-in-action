package main

import "fmt"

var globalX int

func main() {
	localX := 123
	fmt.Println(globalX)
	fmt.Println(localX)
	// Be careful! `:=` can also change the global value
	globalX := 456
	fmt.Println(globalX)

	// Error !You can not REDEFINE one variable (but two is ok!, see later)
	// localX := 789
	// But of cause, you can change it value
	localX = 987
	fmt.Println(localX)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0											  //
// 123											  //
// 456											  //
// 987											  //
////////////////////////////////////////////////////
