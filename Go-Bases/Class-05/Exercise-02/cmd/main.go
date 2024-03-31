package main

import "class/bases/five/employee/internal"

func main() {
	person := internal.Person{Name: "Julian", ID: 1, DateOfBirth: "01/05/1993"}
	employee := internal.Employee{ID: 4, Position: "Software developer", Person: person}

	internal.Employee.PrintEmployee(employee)
}
