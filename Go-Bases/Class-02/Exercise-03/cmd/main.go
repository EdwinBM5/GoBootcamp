package main

import "fmt"

func main() {
	/*
		Create an application that receives a variable with the number of the month.

		Depending on the number, print the corresponding month in text.
		Can you think of different ways to solve it? Which one would you choose and why?
		Ex: 7, July.

		Note: Validate that the number of the month is correct.
	*/

	var month int = 10
	var months map[int]string = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "Agust",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	fmt.Printf("%d, %s\n", month, months[month])
}
