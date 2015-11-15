package main

import "fmt"

func main() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}

	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [1 2 0]										  //
// [1 2 3 4]									  //
// [0 0 100 0 200]								  //
// [{user1 10} {user2 20}]						  //
////////////////////////////////////////////////////
