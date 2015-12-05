package main

import "fmt"

type ReadCloser interface {
	Close() error
}

type Response struct {
	Body ReadCloser
}

func main() {

}
