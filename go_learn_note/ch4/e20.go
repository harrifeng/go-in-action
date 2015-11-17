package main

import "fmt"

type File struct {
	name string
	size int
	attr struct {
		perm int
		ower int
	}
}

func main() {
	f := File{
		name: "test.txt",
		size: 1025,
	}

	var attr = struct {
		perm int
		ower int
	}{2, 0755}

	fmt.Println(f)
	f.attr = attr
	fmt.Println(f)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// {test.txt 1025 {0 0}}						  //
// {test.txt 1025 {2 493}}						  //
////////////////////////////////////////////////////
