package main

import "fmt"

/*
An animal shelter needs to calculate how much food to buy for pets.
At the moment they only have tarantulas, hamsters, dogs and cats, but they hope to be able to shelter many more animals.

They have the following data:

Dog 10 kg of food.
Cat 5 kg of food.
Hamster 250 g of food.
Tarantula 150 g of food.
You are requested to:

Implement an Animal function that receives as a parameter a text type value with the specified animal and
return a function and a message (in case the animal does not exist).

A function for each animal that calculates the amount of food based on the amount of the specified animal type.

const (
	dog    = "dog"
	cat    = "cat"
)

...

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

...

var amount float64

amount += animalDog(10)
amount += animalCat(10)
*/

func CalculateDogFood(amount int) float64 {
	return float64(amount) * 10
}

func CalculateCatFood(amount int) float64 {
	return float64(amount) * 5
}
func CalculateHamsterFood(amount int) float64 {
	return float64(amount) * 0.25
}
func CalculateTarantulaFood(amount int) float64 {
	return float64(amount) * 0.15
}

func animal(animal string) (func(int) float64, string) {
	switch animal {
	case "dog":
		return CalculateDogFood, ""
	case "cat":
		return CalculateCatFood, ""
	case "hamster":
		return CalculateHamsterFood, ""
	case "tarantula":
		return CalculateTarantulaFood, ""
	default:
		return nil, "Error: animal does not exist"
	}
}

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {

	animalDog, _ := animal(dog)
	animalCat, _ := animal(cat)
	animalHamster, _ := animal(hamster)
	animalTarantula, _ := animal(tarantula)

	var amount float64

	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(10)
	amount += animalTarantula(10)

	fmt.Printf("\nThe amount of food is: %.2f kg\n", amount)

}
