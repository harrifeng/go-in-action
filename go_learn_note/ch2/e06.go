package main

import "fmt"

func main() {
	// slice here is reference type, no copy happened when ranging
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
		}

		fmt.Println(i, v)
	}
	fmt.Println(s)
}
