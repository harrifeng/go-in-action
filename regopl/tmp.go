package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary int
}

var emp = []Employee{
	{1, "Adaa", 2500},
	{2, "Bell", 3500},
}

func EmployeeByID(id int) *Employee {
	if id == 1 {
		return &emp[0]
	}
	if id == 2 {
		return &emp[1]
	}
	return nil
}

func main() {
	EmployeeByID(1).Salary = 24
	fmt.Println(EmployeeByID(1).Salary)
}
