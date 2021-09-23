package main

import "fmt"

func main() {
	/*
		Variable declaration:
			1) Shorthand declaration (declaration + assignment):
				This cannot be used outside function bodies.
				example: x := 42
			2) Using the var keyword:
				This can be used outside function bodies.
				example: var x = 42
						 var z int = 75
	*/
	fmt.Println("Hello There!")
	x := 45 // Short hand declaration of variables
	fmt.Println(x + 55)
	var z string = "This is a string!" // The var keyword in action
	fmt.Println(z)
}
