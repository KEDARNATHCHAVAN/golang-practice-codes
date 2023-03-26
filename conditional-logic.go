package main

import (
	"fmt"
)

func main() {
	var num int
	var res string
	fmt.Println("Enter an integer:")
	fmt.Scan(&num)

	//if, else syntax similar to C and C++
	if num < 0 {
		res = "negative"
	} else if num == 0 {
		res = "zero"
	} else {
		res = "positive"
	}
	fmt.Println("The entered number is", res)

	// we can also initialise a variable in if condition
	if num = num * (-1); num < 0 {
		res = "positive"
	} else if num == 0 {
		res = "zero"
	} else {
		res = "negative"
	}
	fmt.Println("The entered number is", res)
}
