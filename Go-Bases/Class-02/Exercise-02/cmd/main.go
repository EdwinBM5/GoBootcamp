package main

import "fmt"

func main() {
	/*
		A bank wants to grant loans to its customers, but not everyone can access them. For this reason,
		it has certain rules to know which clients can be granted loans. It only grants loans to customers
		who are over 22 years of age, are employed and have been working for more than one year. Within the loans it grants,
		it will not charge interest to those who have a salary of more than $100,000.

		It is necessary to make an application that receives these variables and prints a message according to each case.

		Tip: your code must be able to print at least 3 different messages.

	*/

	var age int = 23
	var salary int = 100000
	var isEmployed bool = true
	var yearsWorking int = 3

	if age > 22 && isEmployed && yearsWorking > 1 {
		if salary > 100000 {
			fmt.Println("You can get a loan without interest")
		} else {
			fmt.Println("You can get a loan with interest")
		}
	} else {
		fmt.Println("You can't get a loan")
	}
}
