package main

import "log"

// interfaces are a list of things type/s should be able to do.
type Animal interface {
	numberOfLegs() int
	speaks() string
}

type Dog struct {
	Name string
	Age  int
}

type Counter struct {
	Count int
}

func main() {
	dog := Dog{
		Name: "Barry",
		Age:  3,
	}
	dog.speaks()

	value := Counter{
		Count: 0,
	}

	value.IncrementCopy()
	log.Println(value.Count, "copy")
	value.IncrementReal()
	log.Println(value.Count, "real")

}

// here (d Dog) is a reciever.
// A receiver is like a parameter that a method uses to access the data of a struct.
// value receiver - works on copy of a struct
// func (d Dog) speaks() string {
// 	return "Woof"
// }

// pointer receiver works on atual struct
func (d *Dog) speaks() string {
	return "Woof"
}

// Value receiver -> wont change the original count
func (c Counter) IncrementCopy() {
	c.Count++
}

// Pointer receiver
func (c *Counter) IncrementReal() {
	c.Count++
}
