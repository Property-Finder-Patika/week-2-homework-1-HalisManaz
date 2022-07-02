package main

import "fmt"

func main() {
	const (
		_ = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	// First value of iota is zero. After that
	// Thanks to iota feature days of weeks take values 1,2,3,4,5,6,7 orderly

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
