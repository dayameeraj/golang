package main

import "fmt"

func main() {
	// note : error encountered when var is declared but not used.
	fmt.Println("hello world")
	var greet string
	var i int
	i = 64
	greet = "Hello"
	fmt.Println(greet,i)
	greeting := greetUser()

	fmt.Println(greeting)
}


func greetUser() string {
	return "Hey There"
}