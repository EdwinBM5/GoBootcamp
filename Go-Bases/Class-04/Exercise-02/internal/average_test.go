package internal

import "testing"

func TestStudentAverage(t *testing.T) {
	/*
		Exercise 2 - Calculate average
		The school reported that the operations to calculate the average are not being performed correctly,
		so now it is up to us to perform the corresponding tests:

		Calculate the average of the students' grades.
	*/
	t.Run("Average of student notes", func(t *testing.T) {
		//arrange
		notes := []float64{10, 20, 30, 40, 50}
		expectedAverage := 30.0
		//act
		result, err := StudentAverage(notes...)
		//assert
		if err != "" {
			t.Errorf("Expected %v but got %v", expectedAverage, result)
		}
	})
}
