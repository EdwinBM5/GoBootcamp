package internal

import (
	"testing"
)

func TestCalculateSalaryTax(t *testing.T) {
	/*
		Exercise 1 - Testing the salary tax
		The chocolate company that previously needed to calculate the tax of their employees at the time of
		depositing their salary now asked us to validate that the calculations of these taxes are correct. For this they asked us to perform the corresponding tests for:

		Calculate the tax in case the employee earns below $50,000.
		Calculate the tax in case the employee earns more than $50,000.
		Calculate the tax in case the employee earns more than $150,000.
	*/
	t.Run("Tax on salaries below $50.000", func(t *testing.T) {
		//arrange
		salary := 40000.0
		expectedSalaryTax := 40000.0
		//act
		result := CalculateSalaryTax(salary)

		//assert
		if result != expectedSalaryTax {
			t.Errorf("Expected %v but got %v", expectedSalaryTax, result)
		}
	})

	t.Run("Tax on salaries above $50.000", func(t *testing.T) {
		//arrange
		salary := 60000.0
		expectedSalaryTax := salary - salary*0.17

		//act
		result := CalculateSalaryTax(salary)

		//assert
		if result != expectedSalaryTax {
			t.Errorf("Expected %v but got %v", expectedSalaryTax, result)
		}
	})

	t.Run("Tax on salaries above $150.000", func(t *testing.T) {
		//arrange
		salary := 160000.0
		expectedSalaryTax := salary - salary*0.27

		//act
		result := CalculateSalaryTax(salary)

		//assert
		if result != expectedSalaryTax {
			t.Errorf("Expected %v but got %v", expectedSalaryTax, result)
		}
	})
}
