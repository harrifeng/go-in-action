package main

import "fmt"

type HF func(string) string

// With following code, compiler will complains:
// cannot convert help (type func(string) string) to type HF
// type HF func(string) int

func help(s string) string {
	return s
}

func tester(ii int, hh HF) {
	fmt.Println(hh("in tester"))
}

func main() {
	tester(23, HF(help))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// in tester									  //
////////////////////////////////////////////////////
