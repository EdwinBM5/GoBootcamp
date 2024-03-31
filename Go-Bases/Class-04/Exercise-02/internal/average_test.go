package internal

import "testing"

func TestStudentAverage(t *testing.T) {
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
