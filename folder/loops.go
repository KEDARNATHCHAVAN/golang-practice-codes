package main

import "fmt"

func main() {
	colors := []string{"red", "blue", "greed"}

	// for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// iterating through the loop
	for i := range colors {
		fmt.Println(colors[i])
	}

	// just like while() loop
	val := 1
	for val < 10 {
		fmt.Println("the value is:", val)
		val++
	}

	// we can use break, continue, and goto also
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("The sum is:", sum)
		if sum > 100 {
			goto end
		}
	}

	//label for goto statement
end:
	fmt.Println("End of program !")
}
