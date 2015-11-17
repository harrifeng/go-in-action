package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
	}

	if v, ok := m["a"]; ok { // m["a"] will return two value!
		fmt.Println(v)
	}

	fmt.Println(m["c"])

	m["b"] = 2
	delete(m, "c")
	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1											  //
// 0											  //
// 2											  //
// a 1											  //
// b 2											  //
////////////////////////////////////////////////////
