package main

//Use log for logging events, errors, and debugging in a structured way.
import "log"

func main() {

	var myString string = "LFG"

	log.Println(myString, "before")
	// & to reference the memory address of the string
	pointerFunc(&myString)
	log.Println(myString, "after")
}

func pointerFunc(s *string) {
	// *  is used to get the value at the given memory address
	newValue := "NO"
	*s = newValue
}
