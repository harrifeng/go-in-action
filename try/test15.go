package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune { return r + 1 }

func main() {

	fmt.Println(strings.Map(add1, "ABC"))
	fmt.Println(strings.Map(add1, "WXY"))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// BCD											  //
// XYZ											  //
////////////////////////////////////////////////////
