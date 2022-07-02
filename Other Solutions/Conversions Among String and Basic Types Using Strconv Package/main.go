package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := "1881"
	// Use Atoi to parse string.
	num, _ := strconv.Atoi(value)

	fmt.Println(num)
	fmt.Println(num + 57)

	x := 700

	// Use Itoa on an int.
	result := strconv.Itoa(x)
	fmt.Println(result)

	// strconv quote usage
	q := strconv.Quote("Double quotes still remaining thanks to strconv.queote, also世界")
	fmt.Println(q)
}
