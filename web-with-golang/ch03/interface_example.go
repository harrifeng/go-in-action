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

//Human method SayHi
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s`\n", h.name, h.phone)
}

//Human method Sing
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human method Guzzle
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//Employee method Sayhi
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

//Student BorrowMoney
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

//Employee SpendSalary
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
