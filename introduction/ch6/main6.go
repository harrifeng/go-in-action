package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["B"] = "Boron"
	elements["F"] = "Fluorine"

	fmt.Println(elements["F"])

	if name, ok := elements["A"]; ok {
		fmt.Println(name, ok)
	}

	if name, ok := elements["B"]; ok {
		fmt.Println(name, ok)
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Fluorine										  //
// Boron true									  //
////////////////////////////////////////////////////
