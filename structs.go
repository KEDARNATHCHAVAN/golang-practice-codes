package main

import (
	"fmt"
)

//we can declare structs, just like structures and classes in c++
//If the initial letter of name of struct is capital, that means it is exported (public)
//and if the letter is small then it implies that the struct is non exported (private)

//Note: the structs do not support inheritance in Go unlike JAVA and C++

//each value of a struct is known as field

type Dog struct {
	Breed  string //the same rule applies to elements of struct also for the access specifiers
	Weight int
}

func main() {

	//create a variable and initialise the struct just like constructers
	lucky := Dog{"simple", 10}
	fmt.Println(lucky)
	fmt.Printf("%+v\n", lucky) // to see structs field names and values for debugging
	fmt.Printf("Breed is %v and weight is %v\n", lucky.Breed, lucky.Weight)
	lucky.Breed = "complex"
	fmt.Printf("the breed is changed to %v and the weight is %v\n", lucky.Breed, lucky.Weight)
}
