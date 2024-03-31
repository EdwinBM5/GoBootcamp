package internal

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
