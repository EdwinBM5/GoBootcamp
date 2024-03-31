package internal

func CalculateSalaryTax(salary float64) float64 {
	tax := 0.0
	if salary > 50000 {
		tax += 0.17
	}

	if salary > 150000 {
		tax += 0.10
	}

	return salary - salary*tax
}
