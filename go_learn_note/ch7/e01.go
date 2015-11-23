package main

import (
	"fmt"
	"math"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	fmt.Println(id, x)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}

///////////////////////////////////////////////////////
// <===================OUTPUT===================>	 //
// $ go build e01.go 								 //
// $ ls												 //
// e01	e01.go										 //
// $ GOMAXPROCS=1 time -p ./e01						 //
// 1 9223372030412324865							 //
// 0 9223372030412324865							 //
// real         4.05								 //
// user         4.05								 //
// sys          0.00								 //
// $ GOMAXPROCS=2 time -p ./e01						 //
// 0 9223372030412324865							 //
// 1 9223372030412324865							 //
// real         2.05								 //
// user         4.10								 //
// sys          0.00								 //
///////////////////////////////////////////////////////
