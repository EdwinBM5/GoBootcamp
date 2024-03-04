package main

import "fmt"

func main() {
	/*
	   A seafood company needs to calculate the salary of its employees based on the number of hours worked per month and the category.

	   - Category C, their salary is $1,000 per hour.
	   - Category B, their salary is $1,500 per hour, plus 20% of their monthly salary.
	   - Category A, your salary is $3,000 per hour, plus 50% of your monthly salary.

	   We are asked to generate a function that receives by parameter the number of minutes worked per month,
	   the category and returns her salary.
	*/

	salary, err := CalculateSalary(120, "A")
	if err != "" {
		fmt.Println(err)
	}

	fmt.Printf("Employee salary: %.2f\n", salary)
}

func CalculateSalary(minutes float64, category string) (salary float64, err string) {
	switch category {
	case "A":
		salary = 1000 * (minutes / 60)
	case "B":
		salary = 1500 * (minutes / 60)
		salary += salary * 0.20
	case "C":
		salary = 3000 * (minutes / 60)
		salary += salary * 0.5
	default:
		err = "This category does not exist..."
	}

	return
}
