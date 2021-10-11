package main

import "fmt"

func main() {
	/*
		Go is all about type, It's what is called a typed language.
		This function illustrates a few types in go and also defines a few custom types
	*/

	x := 75 // Variable x is of type integer
	fmt.Printf("Variable x is of type: %T and has value %d \n", x, x)

	var z string = "Gopher it" // Variable z is of type string
	fmt.Printf("Variable z is of type: %T and has value '%s'\n", z, z)

	// NOTE: now, z = 75 Will crash because z is expected to be a string but we are feed it an integer

	// CREATING CUSTOM TYPES
	type dog string // A way to alias types

	// A custom data type called a struct for storing data with certain fields. More on this later!
	type IDCard struct{
		name string
		rollNumber string
	}

	var b dog = "Johnny" // Is of type dog
	fmt.Printf("\nVariable b is of type: %T and has value '%s'\n", b, b)

	var sammy IDCard
	sammy.name = "Sameer"
	sammy.rollNumber = "ABC123"
	fmt.Printf("\nVariable sammy is of type: %T. It has\nname: '%s'\nroll: '%s'\n", sammy, sammy.name, sammy.rollNumber)

	// CONVERSION, not CASTING
	// TODO: Explain what casting is and what the difference between conversion and casting is
	z = string(b) // This is a CONVERSION of types not casting.
	fmt.Printf("\nVariable z is of type: %T and has value '%s'\n", z, z)
}
