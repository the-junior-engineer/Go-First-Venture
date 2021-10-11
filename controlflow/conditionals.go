package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	/*
		Conditionals help run bits of code only if a certain condition is met
		Conditional constructs in Go:
			If-Else
			Switch:
				Compares the value of an expression with multiple values and if anything matches the statements in that block are executed before exiting the switch block
	*/
	rand.Seed(time.Now().UnixNano()) // Ignore this bit

	// If - Else
	// To print out whether a number is even or odd we can use if else statements
	j := rand.Int() // Get a random integer

	if j % 2 == 0 {
		fmt.Printf("The number %d is even\n", j)
	} else {
		fmt.Printf("The number %d is odd\n", j)
	}

	// The above could also be achieved with a switch
	j = rand.Int()
	switch j % 2 {
	case  1 :
		fmt.Printf("The number %d is odd\n", j)
	case 0 :
		fmt.Printf("The number %d is even\n", j)
	// or
	/*
	default:
		fmt.Printf("The number %d is even\n", j)
	*/
	}
}
