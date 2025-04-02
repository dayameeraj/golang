package main

import "log"

func main() {
	// if else
	if 2 > 1 {
		log.Println("here")
	}
	var isTrue bool
	isTrue = true
	animal := "Dog"
	if !isTrue {
		log.Println("True")
	} else {
		log.Println("False")
	}
	if animal == "Dog" && isTrue {
		log.Println("animal is", animal)
	} else if animal == "Dog" && !isTrue {
		log.Println("animal False")
	} else {
		log.Println("else case")
	}

	// switch
	animals := "Dog"
	switch animals {
	case "Dog":
		log.Println("switch case 1", animals)
	case "Cat":
		log.Println("switch case 2", animals)
	default:
		log.Println("None")
	}

	// Loops

	for i := 0; i < 5; i++ {
		log.Println("Loop value", i)
	}

	newAnimals := []string{"dog", "cat", "horse", "zebra", "lion"}

	// _ indetifier ignores first value
	// first value is index
	for _, animal := range newAnimals {
		log.Println("Animal value", animal)
	}

	var text = "Read Dead Redemption"
	// strings are slices of bytes and immutable
	// learn about conversion to []byte and []rune
	log.Println("text[2]", text[2], "subString", text[0:0])
	for i, t := range text {
		log.Println(i, t, "text")
	}

	// looping over custom types

	type UserData struct {
		FirstName string
		LastName  string
	}

	var users []UserData

	users = append(users, UserData{"Dayameeraj", "Biradar"})
	users = append(users, UserData{"John", "Smith"})
	users = append(users, UserData{"Joe", "mama"})

	for i, user := range users {
		log.Println("User", i, user.FirstName, user.LastName)
	}
}
