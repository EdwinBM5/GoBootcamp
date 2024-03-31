package internal

import (
	"testing"
)

func TestCalculateSalaryTax(t *testing.T) {
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
