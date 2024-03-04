package main

import "fmt"

func main() {

	/*
		Fix the following code:

		var 1firstName string
		var lastName string
		var int age
		1lastName := 6
		var driver_license = true
		var person height int
		childsNumber := 2
	*/

	var firstName string = "Juan"
	var lastName string
	var age int = 22
	lastName = "Perez"
	var driverLicense = true
	var personHeight int
	childsNumber := 2

	fmt.Printf("Firstname: %s\nLastname: %s\nAge: %d\nDriver license: %t\nPerson height: %d\nChilds number: %d\n", firstName, lastName, age, driverLicense, personHeight, childsNumber)
}
