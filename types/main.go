package main

import (
	"log"
	"sort"
)

type User struct {
	// capitalized letters are used so that it can accessed outside this package
	// this is how go handles public/ private methods/variables/functions
	FirstName string
	LastName  string
	UserID    string
}

func main() {
	var arthur = User{
		FirstName: "Arthur",
		LastName:  "Morgan",
		UserID:    "1",
	}
	log.Println(arthur)

	// map
	// maps are immutable
	// values stored in maps are randomised so that it is only accessible via key
	myMap := make(map[string]string)
	myMap["num"] = "1"
	myMap["num2"] = "2"
	log.Println(myMap)

	// maps + structs

	myUser := make(map[string]User)

	myUser["daya"] = User{
		FirstName: "Dayameeraj",
		LastName:  "Biradar",
		UserID:    "2",
	}

	log.Println(myUser, "myUser")

	// Array / slices

	var names []string
	names = append(names, "Daya")
	log.Println(names, "names Array")

	var nums []int
	nums = append(nums, 1)
	nums = append(nums, 5)
	nums = append(nums, 2)
	log.Println(nums, "nums before")
	sort.Ints(nums)

	log.Println(nums, "nums after")
	log.Println(nums[0:2], "nums slice")

}
