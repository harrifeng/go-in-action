package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human's phone
}

type Employee struct {
	Human
	speciality string
	phone      string // Employee's phone
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "222-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// If we want to visit Human's phone
	fmt.Println("Bob's person phone is :", Bob.Human.phone)
}
