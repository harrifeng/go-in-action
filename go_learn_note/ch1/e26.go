package main

func main() {

	var a struct {
		x int `a`
	}
	var b struct {
		x int `ab`
	}

	_ = a
	_ = b
	// Error! cannot use a (type struct { x int "a" })
	//as type struct { x int "ab" } in assignment
	// b = a
}
