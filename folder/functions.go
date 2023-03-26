package main

import (
	"fmt"
)

func main() {
	dosomething()
	fmt.Println("the sum is", add(3, 5))
	sum, length := addall(3, 54, 5, 6, 33, 23, 4)
	fmt.Println("the of values is", sum, "and the number of values are", length)
}

// no arguments
func dosomething() {
	fmt.Println("Doing something !")
}

// with arguments and return values
func add(a, b int) int { //the (a,b int) is same as (a int, b int), no need to declare separatly for same datatype arguments
	return a + b
}

// to receive multiple arguments of same type
// and to return multiple values
func addall(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
