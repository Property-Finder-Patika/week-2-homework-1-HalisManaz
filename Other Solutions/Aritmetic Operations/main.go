package main

import "fmt"

func main() {

	var x int = 41
	var y int = 5
	var z int

	z = x + y
	fmt.Printf("Summation --> z: %d\n", z)

	z = x - y
	fmt.Printf("Substraction --> z: %d\n", z)

	z = x * y
	fmt.Printf("Multiplying --> z: %d\n", z)

	z = x / y
	fmt.Printf("Dividing --> z: %d\n", z)

	z = x % y
	fmt.Printf("Remaining --> z: %d\n", z)

	x++
	fmt.Printf("Adding --> z: %d\n", x)

	x--
	fmt.Printf("Decreasing --> %d\n", x)
}
