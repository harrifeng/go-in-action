package main

import (
	"os"
	"text/template"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	t, err := template.New("person").Parse("{{.FirstName}} {{.LastName}} is {{.Age}} years Old")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, Person{
		FirstName: "John",
		LastName:  "Smith",
		Age:       22,
	})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// John Smith is 22 years Old					  //
////////////////////////////////////////////////////
