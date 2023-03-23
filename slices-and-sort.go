package main

import (
	"fmt"
	"sort"
)

func main() {

	//slice
	var colors = []string{"red", "blue", "green"}
	fmt.Println(colors)

	//to add elements
	colors = append(colors, "yellow", "orange")
	fmt.Println(colors)

	//to remove elements
	colors = append(colors[1:]) //removes the 0th element, by starting at 1 till end
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1]) //removes the last element, starting by 0th
	fmt.Println(colors)

	//we can also use make() to create an slice
	numbers := make([]int, 5)
	numbers[0] = 3
	numbers[1] = 5
	numbers[2] = 6
	numbers[3] = 2
	numbers[4] = 55
	fmt.Println("The numbers are", numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
