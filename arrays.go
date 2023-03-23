package main

import "fmt"

func main() {

	//array declaration

	var colors [3]string
	colors[0] = "blue"
	colors[1] = "red"
	colors[2] = "green"
	fmt.Println(colors)
	fmt.Println(colors[2])

	//array initialization

	var numbers = [5]int{2, 3, 54, 6, 5}
	fmt.Println(numbers)
	fmt.Println("The length of array of number is", len(numbers))
}
