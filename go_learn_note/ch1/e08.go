package main

import "fmt"

const x, y int = 1, 2
const s = "Hello, World!"

const (
	a, b      = 10, 100
	c    bool = false
)

func main() {
	// one day const, always const
	const x = "xxx"

	// only x is used
	// unused const variable will not incur Error
	fmt.Println(x)
}
