// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Undeclarables
//
//  1. Declare the variables below:
//      3speed
//      !speed
//      spe?ed
//      var
//      func
//      package
//
//  2. Observe the error messages
//
// NOTE
//  The types of the variables are not important.
// ---------------------------------------------------------

func main() {
	var 3speed int
	var !speed int
	var spe?ed int
	var var int
	var func int
	var package int

	// Error occured:
	// # command-line-arguments
	// .\main.go:29:6: syntax error: unexpected literal 3, expecting name
	// .\main.go:30:6: syntax error: unexpected !, expecting name
	// .\main.go:31:9: invalid character U+003F '?'
	// .\main.go:31:13: syntax error: unexpected int at end of statement
	// .\main.go:32:6: syntax error: unexpected var, expecting name
	// .\main.go:33:6: syntax error: unexpected func, expecting name
	// .\main.go:34:6: syntax error: unexpected package, expecting name
	// PS C:\Users\user\Propery Finder Go Bootcamp\HW\HW2\06-variables\
}
