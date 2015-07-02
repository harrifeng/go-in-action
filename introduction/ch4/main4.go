package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Enter a number: 								  //
// 232343.434									  //
// 464686.868									  //
////////////////////////////////////////////////////
