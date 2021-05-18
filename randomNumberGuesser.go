// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program picks a random number between 1 and 6 for the user to guess

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This function picks a random number between 1 and 6 for the user to guess
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	var randomNumber = rand.Intn(max-min+1) + min
	var userGuess int

	// input
	fmt.Println("This program picks random a number between 1 and 6 for the user to guess")
	fmt.Println()
	fmt.Print("Enter your guess: ")
	fmt.Scanln(&userGuess)
	fmt.Println()

	// process
	if userGuess == randomNumber {
		// output
		fmt.Println("That is correct")
	}
	if userGuess != randomNumber {
		// output
		fmt.Println("Sorry that is incorrect. The correct answer is", randomNumber, ".")
	}

}
