package main

import "fmt"

func main() {
	// note : error encountered when var is declared but not used.
	fmt.Println("hello world")
	var greet string
	var i int = 64
	greet = "Hello"
	fmt.Println(greet, i)
	// declaring and initialising variable , type is assigned based on the value
	greeting := greetUser()
	first, sec := multiReturn()
	fmt.Println(greeting, first, sec)
}

func greetUser() string {
	return "Hey There"
}

func multiReturn() (string, int) {
	return "goLang", 19
}
