package internal

import "fmt"

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("ID: %d, Position: %s, ID Person: %d, Name: %s, Date of birth: %s\n", e.ID, e.Position, e.ID, e.Name, e.DateOfBirth)
}
