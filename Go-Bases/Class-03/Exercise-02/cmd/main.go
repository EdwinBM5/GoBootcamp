package main

import "fmt"

func main() {
	/*
		A school needs to calculate the average (per student) of its grades.
		It is requested to generate a function in which it can be passed N integers and return the average.
		Negative grades cannot be entered.
	*/

	average, err := StudentAverage(10, 20, 30, 40, 50)
	if err != "" {
		fmt.Println(err)
	}

	fmt.Printf("Student average: %.2f\n", average)
}

func StudentAverage(notes ...float64) (average float64, err string) {
	var totalNotes int = len(notes)
	for _, note := range notes {
		if note < 0 {
			err = "Negative notes found, please verify..."
			return
		}

		average += note
	}

	average /= float64(totalNotes)

	return
}
