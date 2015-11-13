package main

func except() {
	recover()
}

func test() {
	defer except()
	panic("test panic")
}

func main() {
	test()
}
