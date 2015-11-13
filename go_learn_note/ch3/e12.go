package main

import (
	"fmt"
	"os"
)

func test() error {
	f, err := os.Create("test.txt")

	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString("Hello World")
	return nil
}

func main() {
	err := test()
	if err != nil {
		fmt.Println("Error")
	}
}
