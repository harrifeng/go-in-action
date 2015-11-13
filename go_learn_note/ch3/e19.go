package main

func test() {
	defer func() {
		recover()
	}()
	panic("test panic")
}

func main() {
	test()
}
