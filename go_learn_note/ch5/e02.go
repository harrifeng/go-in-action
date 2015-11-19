package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() {
	// Pass by Value, &self is changing
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() {
	// Pass by reference, self is always the same
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()

}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// Data: 0x82024c220							  //
// Value: 0x82024c240							  //
// Pointer: 0x82024c220							  //
// Value: 0x82024c250							  //
// Pointer: 0x82024c220							  //
////////////////////////////////////////////////////
