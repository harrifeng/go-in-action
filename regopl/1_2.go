package main

import (
	"fmt"
	"os"
)

func main() {
	for i, c := range os.Args {
		fmt.Println(i, c)
	}
}
