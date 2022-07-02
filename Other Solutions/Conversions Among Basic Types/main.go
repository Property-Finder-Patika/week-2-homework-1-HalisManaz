package main

import "fmt"

func main() {
	// int to float conversation
	var x int = 19
	var y int = 21

	var z float64

	z = float64(x) * float64(y)

	z++

	// float to int conversation
	t := 5 / 2
	fmt.Println(int(t))
}
