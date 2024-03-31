package main

import (
	"class/bases/four/tax/internal"
	"fmt"
)

func main() {
	/*
		A chocolate company needs to calculate the tax of its employees at the time of depositing the salary,
		to meet the objective it is necessary to create a function that returns the tax of a salary.

		Taking into account that if the person earns more than $50,000, 17% will be deducted from the salary and
		if the person earns more than $150,000, 10% will be deducted in addition (27% in total).
	*/

	salary := 160000.0
	fmt.Printf("Employee salary: %.2f\nSalary with taxes: %.2f\n", salary, internal.CalculateSalaryTax(salary))
}
