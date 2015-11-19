package main

import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 10)}
}

func (*Queue) Push(e interface{}) error {
	panic("not implemented")
}

func (self *Queue) length() int {
	return len(self.elements)
}

func main() {
	q := NewQueue()
	fmt.Printf("%d\n", q.length())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 10											  //
////////////////////////////////////////////////////
