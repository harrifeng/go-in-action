package main

import "fmt"

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	xPlusy, xTimesy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPlusy)
	fmt.Printf("%d * %d = %d\n", x, y, xTimesy)
}
