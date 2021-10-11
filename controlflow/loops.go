package main

import "fmt"

func main(){
	/*
		Loops are used to repeat an action a certain number of times
		Looping construct in Go is the for loop
		Note that Go's "while" loop is "for" https://tour.golang.org/flowcontrol/3
	*/

	// For loop demonstration
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the for loop iteration number: %d\n", i + 1)
	}
	fmt.Println("======== End of For loop ========")
	// While loop demonstration
	j := 0
	for j < 5 {
		fmt.Printf("This is the while loop iteration number: %d\n", j + 1)
		j++
	}

	// Infinite for loop
	for {

	}
}



