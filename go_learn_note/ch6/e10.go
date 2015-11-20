package main

import "fmt"

type Tester interface {
	Do()
}

// Regard function AS type
type FuncDo func()

// the function implement the interface!
// then call itself!
func (self FuncDo) Do() { self() }

func main() {
	// t is Tester type
	// FuncDo implement the Do(), it is also the Tester type
	var t Tester = FuncDo(func() { fmt.Println("Hello World") })
	t.Do()
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Hello World									  //
////////////////////////////////////////////////////
