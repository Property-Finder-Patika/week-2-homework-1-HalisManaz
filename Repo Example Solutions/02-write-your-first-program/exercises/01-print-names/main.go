// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Print names
//
//  Print your name and your best friend's name using
//  Println twice
//
// EXPECTED OUTPUT
//  YourName
//  YourBestFriendName
//
// BONUS
//  Use `go run` first.
//  And after that use `go build` and run your program.
// ---------------------------------------------------------

import "fmt"

func main() {
	// Print my name and best friend
	fmt.Println("Halis")
	fmt.Println("Zeynep")

	// Use Print instead of Println
	fmt.Print("Halis")

	// Seperate names with comma
	fmt.Println("Halis" + "," + "Zeynep")

	// fmt.Println(Halis)
	// Error occured while removing double quote from string
	// .\main.go:39:14: undefined: Halis

	// Error occured while package and import statement at bottom of the file
	// main.go:28:1: expected 'package', found 'func'

}
