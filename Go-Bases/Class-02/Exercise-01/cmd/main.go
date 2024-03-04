package main

import (
	"fmt"
)

func main() {
	/*
		The Royal Spanish Academy wants to know how many letters a word has and then have each letter spelled separately.
		For that they will have to:

		Create an application that receives a variable with the word and print the number of letters contained in the word.
		Then, print each of the letters.
	*/

	var word string = "Hello"

	fmt.Printf("The word %s has %d letters\n", word, len(word))

	for index, letter := range word {
		fmt.Printf("%d. %c\n", index, letter)
	}
}
