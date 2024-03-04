package main

import "fmt"

func main() {
	/*
		A meteorology company wants an application where it can have the temperature, humidity,
		and atmospheric pressure from different places.

		Declare 3 variables specifying the data type, as value they should have the temperature,
		humidity, and pressure from where you are. Print the values of the variables in the console.
		What data type would you assign to the variables?
	*/
	var temperature int
	var humidity float64
	var pressure float64

	temperature = 25
	humidity = 0.5
	pressure = 1013.25

	fmt.Printf("Temperature: %d\nHumidity: %f\nPressure: %f\n", temperature, humidity, pressure)
}
