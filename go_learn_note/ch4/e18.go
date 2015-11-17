package main

import "fmt"

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}

	fmt.Println(n1)
	fmt.Println(n2)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// {0 1 <nil> <nil>}							  //
// {0 2 <nil> 0x82024a020}						  //
////////////////////////////////////////////////////
