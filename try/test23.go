package main

import "fmt"
import "os"
import "bufio"

func main() {
	fmt.Println("hello")
	outfile, _ := os.Create("tmp.txt")
	writer := bufio.NewWriter(outfile)
	writer.WriteString("hello, tmp")
	writer.Flush()
}
