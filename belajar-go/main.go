package main

import "belajar-go/helpers"

func main() {
	helpers.Greet()

	// helpers.greet()

	var person = helpers.Person{}

	person.Invokegreet()
}