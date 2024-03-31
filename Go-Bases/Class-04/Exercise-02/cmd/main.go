package main

import (
	"class/bases/four/average/internal"
	"fmt"
)

func main() {
	/*
		A school needs to calculate the average (per student) of its grades.
		It is requested to generate a function in which it can be passed N integers and return the average.
		Negative grades cannot be entered.
	*/

	average, err := internal.StudentAverage(10, 20, 30, 40, 50)
	if err != "" {
		fmt.Println(err)
	}

	fmt.Printf("Student average: %.2f\n", average)
}
