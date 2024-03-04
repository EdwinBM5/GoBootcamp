package main

import "fmt"

func main() {
	/*
		A programming student tried to make variable declarations with their respective types in Go,
		but had several doubts while doing it. From this, he provided us with his code and asked
		for the help of an experienced developer who can:

		Verify his code and make the necessary corrections.

		var lastName string = "Smith"
		var age int = "35"
		boolean := "false"
		var salary string = 45857.90
		var firstName string = "Mary"
	*/

	var lastName string = "Smith"
	var age int = 35
	boolean := false
	var salary float32 = 45857.90
	var firstName string = "Mary"

	fmt.Printf("Lastname: %s\nAge: %d\nBoolean: %t\nSalary: %f\nFirstname: %s\n", lastName, age, boolean, salary, firstName)
}
