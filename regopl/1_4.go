package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filep := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filep)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filep)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("str: '%s', times: %d, pos: %s\n", line, n, filep[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filep map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filep[input.Text()] += fmt.Sprintf("%v ", f.Name())
	}
}

/////////////////////////////////////////////////////
// <===================OUTPUT===================>  //
// $ printf "%s\n" a b c cd b a b | go run 1_4.go  //
// a 2											   //
// b 3											   //
/////////////////////////////////////////////////////
