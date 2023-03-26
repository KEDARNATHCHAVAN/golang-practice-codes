package main

import (
	"fmt"
	"math" //other example : math.pi
)

func main() {

	i1, i2, i3 := 12, 45, 68
	sum := i1 + i2 + i3
	fmt.Println("The sum is", sum)

	f1, f2, f3 := 43.54, 65.43, 65.66
	floatsum := f1 + f2 + f3
	fmt.Println("The float sum is", floatsum)

	floatround := math.Round(floatsum)
	fmt.Println("The rounded value is", floatround)
}
