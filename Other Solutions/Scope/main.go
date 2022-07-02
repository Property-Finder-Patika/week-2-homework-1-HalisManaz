package main

import "fmt"

// check() function is the below and call it on the same file
// However greet() function in greet.go file. Thanks to same package, greet function could be callable. Therefore, this is a scope example in Golang

func main() {
	greet()
	check()
}

func check() {
	fmt.Println("How are you today?")
}
