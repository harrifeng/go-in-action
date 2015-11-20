package main

import "fmt"

type Data struct {
	id int
}

func (self Data) String() string {
	return ""
}

func main() {
	// fmt.Stringer is interface type, _ is variable, but we don't
	// care about it.
	var _ fmt.Stringer = (*Data)(nil)

	// IF String() is not implemented =>
	// *Data does not implement fmt.Stringer (missing String method)
}
