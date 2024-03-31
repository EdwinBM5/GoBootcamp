package main

import "fmt"

func main() {
	/*
		Professors at a university in Colombia need to calculate some grading statistics for students in a course.
		They need to calculate the minimum, maximum and average values of their grades.

		For that, it is requested to generate a function that indicates what type of calculation they want
		to perform (minimum, maximum or average) and that returns another function and a message (in case the calculation is not defined)
		that can be passed a number N of integers and returns the calculation that was indicated in the previous function.

		Example:

		const (

			minimum = "minimum"
			average = "average"
			maximum = "maximum"
		)


		...

		minFunc, err := operation(minimum)
		averageFunc, err := operation(average)
		maxFunc, err := operation(maximum)

		...

		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	*/

	minFunc, err := Operation(Minimum)
	if err != "" {
		fmt.Println(err)

	}

	averageFunc, err := Operation(Average)
	if err != "" {
		fmt.Println(err)

	}

	maxFunc, err := Operation(Maximum)
	if err != "" {
		fmt.Println(err)

	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5, 1)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Minimum value: %d\n", minValue)
	fmt.Printf("Average value: %d\n", averageValue)
	fmt.Printf("Maximum value: %d\n", maxValue)
}

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func MinFunc(numbers ...int) (minNum int) {
	minNum = numbers[0]
	for _, num := range numbers {
		if num < minNum {
			minNum = num
		}
	}

	return
}

func MaxFunc(numbers ...int) (maxNum int) {
	maxNum = numbers[0]
	for _, num := range numbers {
		if num > maxNum {
			maxNum = num
		}
	}

	return
}

func AverageFunc(numbers ...int) (avgNum int) {
	var lenNumbers = len(numbers)
	for _, num := range numbers {
		avgNum += num
	}

	avgNum /= lenNumbers

	return
}

func Operation(opr string) (fn func(nums ...int) int, err string) {
	switch opr {
	case Minimum:
		fn = MinFunc
	case Maximum:
		fn = MaxFunc
	case Average:
		fn = AverageFunc
	default:
		err = "This operation is invalid..."
	}

	return
}
