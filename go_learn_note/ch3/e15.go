package main

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // err is interface{} type and now convert it to string
		}
	}()

	panic("panic error!")
}

func main() {
	test()
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// panic error!									  //
////////////////////////////////////////////////////
