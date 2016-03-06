package pkg

import "fmt"

type FHR struct {
	Wwords []uint64
	words  []uint64
}

type QT []uint64

var Wwords = "workds, hello"

func (QT) Display() {
	fmt.Println("QT")
}

func (FHR) Display() {
	fmt.Println("FHR")
}
