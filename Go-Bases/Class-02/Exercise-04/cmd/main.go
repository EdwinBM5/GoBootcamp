package main

import "fmt"

func main() {
	/*
		An employee of a company wants to know the name and age of one of his employees. According to the map below,
		help print Benjamin's age.

		var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

		On the other hand it is also necessary:

		- Know how many of your employees are over the age of 21.
		- Add a new employee to the list, named Federico, who is 25 years old.
		- Eliminate Pedro from the map.
	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin's age")
	fmt.Printf("Benjamin's age is %d\n", employees["Benjamin"])

	fmt.Println("\nHow many of your employees are over the age of 21")
	for key, value := range employees {
		if value > 21 {
			fmt.Printf("%s is %d so is over 21\n", key, value)
		}
	}

	fmt.Println("\nAdd a new employee to the list, named Federico, who is 25 years old")
	employees["Federico"] = 25

	fmt.Println("\nEliminate Pedro from the map")
	delete(employees, "Pedro")

	fmt.Println("\nPrint the employees map")
	for key, value := range employees {
		fmt.Printf("%s is %d\n", key, value)
	}
}
