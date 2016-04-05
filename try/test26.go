package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr := net.ParseIP(name)
	fmt.Printf("%T\n", addr)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
		fmt.Println("Default Mask is", addr.DefaultMask())
	}
	os.Exit(0)
}
