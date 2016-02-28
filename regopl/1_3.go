package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	s, sep := "", ""
	for _, c := range os.Args {
		s += sep + c
		sep = ""
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
