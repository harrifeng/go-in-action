package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Human SayHi method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human Sing method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la..", lyrics)
}

// Employee overload SayHi
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

// Student and Employee all have SayHi and Sing methods
// and thus they are all implements the interface Men
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd", 5000}

	var i Men

	// i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i can also store Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice Men
	fmt.Println("Let's use a slice of Men and see what happends")
	x := make([]Men, 3)

	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
